package implements

type CidadeUF struct {
	City  string
	State string
}

func IsCoastalCities(city, state string) bool {
	coastalCities := GetCoastalCities()
	for _, coastalCitite := range coastalCities {
		if coastalCitite.City == city && coastalCitite.State == state {
			return true
		}
	}
	return false
}

// https://www.ibge.gov.br/geociencias/organizacao-do-territorio/estrutura-territorial/24072-municipios-defrontantes-com-o-mar.html
func GetCoastalCities() []CidadeUF {
	return []CidadeUF{
		{
			City:  "Augusto Corrêa",
			State: "PA",
		},
		{
			City:  "Bragança",
			State: "PA",
		},
		{
			City:  "Chaves",
			State: "PA",
		},
		{
			City:  "Colares",
			State: "PA",
		},
		{
			City:  "Curuçá",
			State: "PA",
		},
		{
			City:  "Magalhães Barata",
			State: "PA",
		},
		{
			City:  "Maracanã",
			State: "PA",
		},
		{
			City:  "Marapanim",
			State: "PA",
		},
		{
			City:  "Quatipuru",
			State: "PA",
		},
		{
			City:  "Salinópolis",
			State: "PA",
		},
		{
			City:  "Salvaterra",
			State: "PA",
		},
		{
			City:  "São Caetano de Odivelas",
			State: "PA",
		},
		{
			City:  "São João de Pirabas",
			State: "PA",
		},
		{
			City:  "Soure",
			State: "PA",
		},
		{
			City:  "Tracuateua",
			State: "PA",
		},
		{
			City:  "Vigia",
			State: "PA",
		},
		{
			City:  "Viseu",
			State: "PA",
		},
		{
			City:  "Amapá",
			State: "AP",
		},
		{
			City:  "Calçoene",
			State: "AP",
		},
		{
			City:  "Macapá",
			State: "AP",
		},
		{
			City:  "Oiapoque",
			State: "AP",
		},
		{
			City:  "Alcântara",
			State: "MA",
		},
		{
			City:  "Apicum-Açu",
			State: "MA",
		},
		{
			City:  "Araioses",
			State: "MA",
		},
		{
			City:  "Axixá",
			State: "MA",
		},
		{
			City:  "Bacabeira",
			State: "MA",
		},
		{
			City:  "Bacuri",
			State: "MA",
		},
		{
			City:  "Bacurituba",
			State: "MA",
		},
		{
			City:  "Barreirinhas",
			State: "MA",
		},
		{
			City:  "Bequimão",
			State: "MA",
		},
		{
			City:  "Cajapió",
			State: "MA",
		},
		{
			City:  "Cândido Mendes",
			State: "MA",
		},
		{
			City:  "Carutapera",
			State: "MA",
		},
		{
			City:  "Cedral",
			State: "MA",
		},
		{
			City:  "Cururupu",
			State: "MA",
		},
		{
			City:  "Godofredo Viana",
			State: "MA",
		},
		{
			City:  "Guimarães",
			State: "MA",
		},
		{
			City:  "Humberto de Campos",
			State: "MA",
		},
		{
			City:  "Icatu",
			State: "MA",
		},
		{
			City:  "Luís Domingues",
			State: "MA",
		},
		{
			City:  "Paço do Lumiar",
			State: "MA",
		},
		{
			City:  "Paulino Neves",
			State: "MA",
		},
		{
			City:  "Porto Rico do Maranhão",
			State: "MA",
		},
		{
			City:  "Primeira Cruz",
			State: "MA",
		},
		{
			City:  "Raposa",
			State: "MA",
		},
		{
			City:  "Rosário",
			State: "MA",
		},
		{
			City:  "Santa Rita",
			State: "MA",
		},
		{
			City:  "Santo Amaro do Maranhão",
			State: "MA",
		},
		{
			City:  "São João Batista",
			State: "MA",
		},
		{
			City:  "São José de Ribamar",
			State: "MA",
		},
		{
			City:  "São Luís",
			State: "MA",
		},
		{
			City:  "Serrano do Maranhão",
			State: "MA",
		},
		{
			City:  "Turiaçu",
			State: "MA",
		},
		{
			City:  "Tutóia",
			State: "MA",
		},
		{
			City:  "Cajueiro da Praia",
			State: "PI",
		},
		{
			City:  "Ilha Grande",
			State: "PI",
		},
		{
			City:  "Luís Correia",
			State: "PI",
		},
		{
			City:  "Parnaíba",
			State: "PI",
		},
		{
			City:  "Acaraú",
			State: "CE",
		},
		{
			City:  "Amontada",
			State: "CE",
		},
		{
			City:  "Aquiraz",
			State: "CE",
		},
		{
			City:  "Aracati",
			State: "CE",
		},
		{
			City:  "Barroquinha",
			State: "CE",
		},
		{
			City:  "Beberibe",
			State: "CE",
		},
		{
			City:  "Camocim",
			State: "CE",
		},
		{
			City:  "Cascavel",
			State: "CE",
		},
		{
			City:  "Caucaia",
			State: "CE",
		},
		{
			City:  "Cruz",
			State: "CE",
		},
		{
			City:  "Fortaleza",
			State: "CE",
		},
		{
			City:  "Fortim",
			State: "CE",
		},
		{
			City:  "Icapuí",
			State: "CE",
		},
		{
			City:  "Itapipoca",
			State: "CE",
		},
		{
			City:  "Itarema",
			State: "CE",
		},
		{
			City:  "Jijoca de Jericoacoara",
			State: "CE",
		},
		{
			City:  "Paracuru",
			State: "CE",
		},
		{
			City:  "Paraipaba",
			State: "CE",
		},
		{
			City:  "São Gonçalo do Amarante",
			State: "CE",
		},
		{
			City:  "Trairi",
			State: "CE",
		},
		{
			City:  "Areia Branca",
			State: "RN",
		},
		{
			City:  "Baía Formosa",
			State: "RN",
		},
		{
			City:  "Caiçara do Norte",
			State: "RN",
		},
		{
			City:  "Canguaretama",
			State: "RN",
		},
		{
			City:  "Ceará-Mirim",
			State: "RN",
		},
		{
			City:  "Parnamirim",
			State: "RN",
		},
		{
			City:  "Extremoz",
			State: "RN",
		},
		{
			City:  "Galinhos",
			State: "RN",
		},
		{
			City:  "Grossos",
			State: "RN",
		},
		{
			City:  "Guamaré",
			State: "RN",
		},
		{
			City:  "Macau",
			State: "RN",
		},
		{
			City:  "Maxaranguape",
			State: "RN",
		},
		{
			City:  "Natal",
			State: "RN",
		},
		{
			City:  "Nísia Floresta",
			State: "RN",
		},
		{
			City:  "Rio do Fogo",
			State: "RN",
		},
		{
			City:  "Pedra Grande",
			State: "RN",
		},
		{
			City:  "Porto do Mangue",
			State: "RN",
		},
		{
			City:  "Tibau",
			State: "RN",
		},
		{
			City:  "São Bento do Norte",
			State: "RN",
		},
		{
			City:  "São Miguel do Gostoso",
			State: "RN",
		},
		{
			City:  "Senador Georgino Avelino",
			State: "RN",
		},
		{
			City:  "Tibau do Sul",
			State: "RN",
		},
		{
			City:  "Touros",
			State: "RN",
		},
		{
			City:  "Baía da Traição",
			State: "PB",
		},
		{
			City:  "Cabedelo",
			State: "PB",
		},
		{
			City:  "Conde",
			State: "PB",
		},
		{
			City:  "João Pessoa",
			State: "PB",
		},
		{
			City:  "Lucena",
			State: "PB",
		},
		{
			City:  "Marcação",
			State: "PB",
		},
		{
			City:  "Mataraca",
			State: "PB",
		},
		{
			City:  "Pitimbu",
			State: "PB",
		},
		{
			City:  "Rio Tinto",
			State: "PB",
		},
		{
			City:  "Santa Rita",
			State: "PB",
		},
		{
			City:  "Barreiros",
			State: "PE",
		},
		{
			City:  "Cabo de Santo Agostinho",
			State: "PE",
		},
		{
			City:  "Fernando de Noronha",
			State: "PE",
		},
		{
			City:  "Goiana",
			State: "PE",
		},
		{
			City:  "Igarassu",
			State: "PE",
		},
		{
			City:  "Ipojuca",
			State: "PE",
		},
		{
			City:  "Ilha de Itamaracá",
			State: "PE",
		},
		{
			City:  "Jaboatão dos Guararapes",
			State: "PE",
		},
		{
			City:  "Olinda",
			State: "PE",
		},
		{
			City:  "Paulista",
			State: "PE",
		},
		{
			City:  "Recife",
			State: "PE",
		},
		{
			City:  "São José da Coroa Grande",
			State: "PE",
		},
		{
			City:  "Sirinhaém",
			State: "PE",
		},
		{
			City:  "Tamandaré",
			State: "PE",
		},
		{
			City:  "Barra de Santo Antônio",
			State: "AL",
		},
		{
			City:  "Barra de São Miguel",
			State: "AL",
		},
		{
			City:  "Coruripe",
			State: "AL",
		},
		{
			City:  "Feliz Deserto",
			State: "AL",
		},
		{
			City:  "Japaratinga",
			State: "AL",
		},
		{
			City:  "Jequiá da Praia",
			State: "AL",
		},
		{
			City:  "Maceió",
			State: "AL",
		},
		{
			City:  "Maragogi",
			State: "AL",
		},
		{
			City:  "Marechal Deodoro",
			State: "AL",
		},
		{
			City:  "Paripueira",
			State: "AL",
		},
		{
			City:  "Passo de Camaragibe",
			State: "AL",
		},
		{
			City:  "Piaçabuçu",
			State: "AL",
		},
		{
			City:  "Porto de Pedras",
			State: "AL",
		},
		{
			City:  "Roteiro",
			State: "AL",
		},
		{
			City:  "São Miguel dos Milagres",
			State: "AL",
		},
		{
			City:  "Aracaju",
			State: "SE",
		},
		{
			City:  "Barra dos Coqueiros",
			State: "SE",
		},
		{
			City:  "Brejo Grande",
			State: "SE",
		},
		{
			City:  "Estância",
			State: "SE",
		},
		{
			City:  "Itaporanga d'Ajuda",
			State: "SE",
		},
		{
			City:  "Pacatuba",
			State: "SE",
		},
		{
			City:  "Pirambu",
			State: "SE",
		},
		{
			City:  "Alcobaça",
			State: "BA",
		},
		{
			City:  "Belmonte",
			State: "BA",
		},
		{
			City:  "Cairu",
			State: "BA",
		},
		{
			City:  "Camaçari",
			State: "BA",
		},
		{
			City:  "Camamu",
			State: "BA",
		},
		{
			City:  "Canavieiras",
			State: "BA",
		},
		{
			City:  "Caravelas",
			State: "BA",
		},
		{
			City:  "Conde",
			State: "BA",
		},
		{
			City:  "Entre Rios",
			State: "BA",
		},
		{
			City:  "Esplanada",
			State: "BA",
		},
		{
			City:  "Igrapiúna",
			State: "BA",
		},
		{
			City:  "Ilhéus",
			State: "BA",
		},
		{
			City:  "Itacaré",
			State: "BA",
		},
		{
			City:  "Ituberá",
			State: "BA",
		},
		{
			City:  "Jaguaripe",
			State: "BA",
		},
		{
			City:  "Jandaíra",
			State: "BA",
		},
		{
			City:  "Lauro de Freitas",
			State: "BA",
		},
		{
			City:  "Maraú",
			State: "BA",
		},
		{
			City:  "Mata de São João",
			State: "BA",
		},
		{
			City:  "Mucuri",
			State: "BA",
		},
		{
			City:  "Nilo Peçanha",
			State: "BA",
		},
		{
			City:  "Nova Viçosa",
			State: "BA",
		},
		{
			City:  "Porto Seguro",
			State: "BA",
		},
		{
			City:  "Prado",
			State: "BA",
		},
		{
			City:  "Salvador",
			State: "BA",
		},
		{
			City:  "Santa Cruz Cabrália",
			State: "BA",
		},
		{
			City:  "Una",
			State: "BA",
		},
		{
			City:  "Uruçuca",
			State: "BA",
		},
		{
			City:  "Valença",
			State: "BA",
		},
		{
			City:  "Vera Cruz",
			State: "BA",
		},
		{
			City:  "Anchieta",
			State: "ES",
		},
		{
			City:  "Aracruz",
			State: "ES",
		},
		{
			City:  "Conceição da Barra",
			State: "ES",
		},
		{
			City:  "Fundão",
			State: "ES",
		},
		{
			City:  "Guarapari",
			State: "ES",
		},
		{
			City:  "Itapemirim",
			State: "ES",
		},
		{
			City:  "Linhares",
			State: "ES",
		},
		{
			City:  "Marataízes",
			State: "ES",
		},
		{
			City:  "Piúma",
			State: "ES",
		},
		{
			City:  "Presidente Kennedy",
			State: "ES",
		},
		{
			City:  "São Mateus",
			State: "ES",
		},
		{
			City:  "Serra",
			State: "ES",
		},
		{
			City:  "Vila Velha",
			State: "ES",
		},
		{
			City:  "Vitória",
			State: "ES",
		},
		{
			City:  "Angra dos Reis",
			State: "RJ",
		},
		{
			City:  "Araruama",
			State: "RJ",
		},
		{
			City:  "Armação dos Búzios",
			State: "RJ",
		},
		{
			City:  "Arraial do Cabo",
			State: "RJ",
		},
		{
			City:  "Cabo Frio",
			State: "RJ",
		},
		{
			City:  "Carapebus",
			State: "RJ",
		},
		{
			City:  "Campos dos Goytacazes",
			State: "RJ",
		},
		{
			City:  "Casimiro de Abreu",
			State: "RJ",
		},
		{
			City:  "Duque de Caxias",
			State: "RJ",
		},
		{
			City:  "Guapimirim",
			State: "RJ",
		},
		{
			City:  "Itaboraí",
			State: "RJ",
		},
		{
			City:  "Itaguaí",
			State: "RJ",
		},
		{
			City:  "Macaé",
			State: "RJ",
		},
		{
			City:  "Magé",
			State: "RJ",
		},
		{
			City:  "Mangaratiba",
			State: "RJ",
		},
		{
			City:  "Maricá",
			State: "RJ",
		},
		{
			City:  "Niterói",
			State: "RJ",
		},
		{
			City:  "Paraty",
			State: "RJ",
		},
		{
			City:  "Quissamã",
			State: "RJ",
		},
		{
			City:  "Rio das Ostras",
			State: "RJ",
		},
		{
			City:  "Rio de Janeiro",
			State: "RJ",
		},
		{
			City:  "São Francisco de Itabapoana",
			State: "RJ",
		},
		{
			City:  "São Gonçalo",
			State: "RJ",
		},
		{
			City:  "São João da Barra",
			State: "RJ",
		},
		{
			City:  "Saquarema",
			State: "RJ",
		},
		{
			City:  "Bertioga",
			State: "SP",
		},
		{
			City:  "Cananéia",
			State: "SP",
		},
		{
			City:  "Caraguatatuba",
			State: "SP",
		},
		{
			City:  "Guarujá",
			State: "SP",
		},
		{
			City:  "Iguape",
			State: "SP",
		},
		{
			City:  "Ilhabela",
			State: "SP",
		},
		{
			City:  "Ilha Comprida",
			State: "SP",
		},
		{
			City:  "Itanhaém",
			State: "SP",
		},
		{
			City:  "Mongaguá",
			State: "SP",
		},
		{
			City:  "Peruíbe",
			State: "SP",
		},
		{
			City:  "Praia Grande",
			State: "SP",
		},
		{
			City:  "Santos",
			State: "SP",
		},
		{
			City:  "São Sebastião",
			State: "SP",
		},
		{
			City:  "São Vicente",
			State: "SP",
		},
		{
			City:  "Ubatuba",
			State: "SP",
		},
		{
			City:  "Guaraqueçaba",
			State: "PR",
		},
		{
			City:  "Guaratuba",
			State: "PR",
		},
		{
			City:  "Matinhos",
			State: "PR",
		},
		{
			City:  "Paranaguá",
			State: "PR",
		},
		{
			City:  "Pontal do Paraná",
			State: "PR",
		},
		{
			City:  "Araquari",
			State: "SC",
		},
		{
			City:  "Araranguá",
			State: "SC",
		},
		{
			City:  "Balneário Arroio do Silva",
			State: "SC",
		},
		{
			City:  "Balneário Camboriú",
			State: "SC",
		},
		{
			City:  "Balneário Barra do Sul",
			State: "SC",
		},
		{
			City:  "Balneário Gaivota",
			State: "SC",
		},
		{
			City:  "Barra Velha",
			State: "SC",
		},
		{
			City:  "Bombinhas",
			State: "SC",
		},
		{
			City:  "Florianópolis",
			State: "SC",
		},
		{
			City:  "Garopaba",
			State: "SC",
		},
		{
			City:  "Governador Celso Ramos",
			State: "SC",
		},
		{
			City:  "Imbituba",
			State: "SC",
		},
		{
			City:  "Itajaí",
			State: "SC",
		},
		{
			City:  "Itapema",
			State: "SC",
		},
		{
			City:  "Itapoá",
			State: "SC",
		},
		{
			City:  "Jaguaruna",
			State: "SC",
		},
		{
			City:  "Laguna",
			State: "SC",
		},
		{
			City:  "Navegantes",
			State: "SC",
		},
		{
			City:  "Palhoça",
			State: "SC",
		},
		{
			City:  "Passo de Torres",
			State: "SC",
		},
		{
			City:  "Paulo Lopes",
			State: "SC",
		},
		{
			City:  "Penha",
			State: "SC",
		},
		{
			City:  "Balneário Piçarras",
			State: "SC",
		},
		{
			City:  "Porto Belo",
			State: "SC",
		},
		{
			City:  "São Francisco do Sul",
			State: "SC",
		},
		{
			City:  "Tijucas",
			State: "SC",
		},
		{
			City:  "Balneário Rincão",
			State: "SC",
		},
		{
			City:  "Lagoa dos Patos",
			State: "RS",
		},
		{
			City:  "Arroio do Sal",
			State: "RS",
		},
		{
			City:  "Balneário Pinhal",
			State: "RS",
		},
		{
			City:  "Capão da Canoa",
			State: "RS",
		},
		{
			City:  "Cidreira",
			State: "RS",
		},
		{
			City:  "Imbé",
			State: "RS",
		},
		{
			City:  "Mostardas",
			State: "RS",
		},
		{
			City:  "Osório",
			State: "RS",
		},
		{
			City:  "Palmares do Sul",
			State: "RS",
		},
		{
			City:  "Rio Grande",
			State: "RS",
		},
		{
			City:  "Santa Vitória do Palmar",
			State: "RS",
		},
		{
			City:  "São José do Norte",
			State: "RS",
		},
		{
			City:  "Tavares",
			State: "RS",
		},
		{
			City:  "Terra de Areia",
			State: "RS",
		},
		{
			City:  "Torres",
			State: "RS",
		},
		{
			City:  "Tramandaí",
			State: "RS",
		},
		{
			City:  "Xangri-lá",
			State: "RS",
		},
	}
}
