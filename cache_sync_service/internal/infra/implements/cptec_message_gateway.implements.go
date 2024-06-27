package implements

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/beriloqueiroz/desafio-back/core/pkg"
	"github.com/sony/gobreaker/v2"
	"golang.org/x/net/html/charset"
	"net/http"
	url2 "net/url"
	"strings"
)

type CptecMessageGateway struct{}

func NewCptecMessageGateway() *CptecMessageGateway {
	initCircuitBreak()
	return &CptecMessageGateway{}
}

type Cidades struct {
	XMLName xml.Name `xml:"cidades" json:"-"`
	Text    string   `xml:",chardata"`
	Cidade  []struct {
		Text string `xml:",chardata"`
		Nome string `xml:"nome"`
		Uf   string `xml:"uf"`
		ID   int64  `xml:"id"`
	} `xml:"cidade"`
}

type PrevisaoCidade struct {
	//XMLName     xml.Name `xml:"cidade" json:"-"`
	Text        string `xml:",chardata"`
	Nome        string `xml:"nome"`
	Uf          string `xml:"uf"`
	Atualizacao string `xml:"atualizacao"`
	Previsao    []struct {
		Text   string `xml:",chardata"`
		Dia    string `xml:"dia"`
		Tempo  string `xml:"tempo"`
		Maxima string `xml:"maxima"`
		Minima string `xml:"minima"`
		Iuv    string `xml:"iuv"`
	} `xml:"previsao"`
}

type PrevisaoOndas struct {
	XMLName     xml.Name `xml:"cidade"`
	Text        string   `xml:",chardata"`
	Nome        string   `xml:"nome"`
	Uf          string   `xml:"uf"`
	Atualizacao string   `xml:"atualizacao"`
	Manha       struct {
		Text     string `xml:",chardata"`
		Dia      string `xml:"dia"`
		Agitacao string `xml:"agitacao"`
		Altura   string `xml:"altura"`
		Direcao  string `xml:"direcao"`
		Vento    string `xml:"vento"`
		VentoDir string `xml:"vento_dir"`
	} `xml:"manha"`
	Tarde struct {
		Text     string `xml:",chardata"`
		Dia      string `xml:"dia"`
		Agitacao string `xml:"agitacao"`
		Altura   string `xml:"altura"`
		Direcao  string `xml:"direcao"`
		Vento    string `xml:"vento"`
		VentoDir string `xml:"vento_dir"`
	} `xml:"tarde"`
	Noite struct {
		Text     string `xml:",chardata"`
		Dia      string `xml:"dia"`
		Agitacao string `xml:"agitacao"`
		Altura   string `xml:"altura"`
		Direcao  string `xml:"direcao"`
		Vento    string `xml:"vento"`
		VentoDir string `xml:"vento_dir"`
	} `xml:"noite"`
}

type Previsao struct {
	Clima PrevisaoCidade
	Ondas PrevisaoOndas
}

func (c *CptecMessageGateway) MessageByLocation(ctx context.Context, city string, state string) (string, error) {
	var result Previsao
	cityId, err := GetCityID(city, state)
	if err != nil {
		return "", err
	}
	climate, err := GetCityClimate(cityId)
	if err != nil {
		return "", err
	}
	result.Clima = climate
	if IsCoastalCities(city, state) {
		waves, err := GetCityWaveForecast(cityId)
		if err != nil {
			return "", err
		}
		result.Ondas = waves
	}
	str, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(str), err
}

func GetCityWaveForecast(id int64) (PrevisaoOndas, error) {
	resp, err := getWithCircuitBreaker(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%d/dia/0/ondas.xml", id))
	if err != nil {
		return PrevisaoOndas{}, err
	}
	defer resp.Body.Close()
	var previsao PrevisaoOndas

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&previsao)
	if err != nil {
		return PrevisaoOndas{}, err
	}
	return previsao, nil
}

func GetCityClimate(id int64) (PrevisaoCidade, error) {
	resp, err := getWithCircuitBreaker(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%d/previsao.xml", id))
	if err != nil {
		return PrevisaoCidade{}, err
	}
	defer resp.Body.Close()
	var previsao PrevisaoCidade

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&previsao)
	if err != nil {
		return PrevisaoCidade{}, err
	}
	return previsao, nil
}

func GetCityID(city string, state string) (int64, error) {
	cityStr, err := pkg.RemoveAccents(city)
	if err != nil {
		return 0, err
	}
	url := "http://servicos.cptec.inpe.br/XML/listaCidades?city=" + url2.PathEscape(cityStr)
	resp, err := getWithCircuitBreaker(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var cidades Cidades
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&cidades)

	if err != nil {
		return 0, err
	}
	var cityId int64 = 0
	for _, cidade := range cidades.Cidade {
		if strings.ToLower(cidade.Uf) == strings.ToLower(state) && strings.ToLower(cidade.Nome) == strings.ToLower(city) {
			cityId = cidade.ID
			break
		}
	}
	if cityId == 0 {
		return 0, errors.New("city not found")
	}
	return cityId, nil
}

func initCircuitBreak() {
	cb = pkg.NewCircuitBreak[*http.Response]("GET CPTEC DATA", 5, 0.5)
}

var cb *gobreaker.CircuitBreaker[*http.Response]

func getWithCircuitBreaker(url string) (*http.Response, error) {
	res, err := cb.Execute(func() (*http.Response, error) {
		return http.DefaultClient.Get(url)
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
