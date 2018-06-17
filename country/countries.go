package country

import (
	"encoding/json"
	"fmt"
	"strings"
)

const JSON = `[
	{
	  "name_pl": "Afganistan",
	  "name_en": "Afghanistan",
	  "code": "AF"
	},
	{
	  "name_pl": "Albania",
	  "name_en": "Albania",
	  "code": "AL"
	},
	{
	  "name_pl": "Algieria",
	  "name_en": "Algeria",
	  "code": "DZ"
	},
	{
	  "name_pl": "Andora",
	  "name_en": "Andorra",
	  "code": "AD"
	},
	{
	  "name_pl": "Angola",
	  "name_en": "Angola",
	  "code": "AO"
	},
	{
	  "name_pl": "Anguilla",
	  "name_en": "Anguilla",
	  "code": "AI"
	},
	{
	  "name_pl": "Antarktyka",
	  "name_en": "Antarctica",
	  "code": "AQ"
	},
	{
	  "name_pl": "Antigua i Barbuda",
	  "name_en": "Antigua and Barbuda",
	  "code": "AG"
	},
	{
	  "name_pl": "Arabia Saudyjska",
	  "name_en": "Saudi Arabia",
	  "code": "SA"
	},
	{
	  "name_pl": "Argentyna",
	  "name_en": "Argentina",
	  "code": "AR"
	},
	{
	  "name_pl": "Armenia",
	  "name_en": "Armenia",
	  "code": "AM"
	},
	{
	  "name_pl": "Aruba",
	  "name_en": "Aruba",
	  "code": "AW"
	},
	{
	  "name_pl": "Australia",
	  "name_en": "Australia",
	  "code": "AU"
	},
	{
	  "name_pl": "Austria",
	  "name_en": "Austria",
	  "code": "AT"
	},
	{
	  "name_pl": "Azerbejdżan",
	  "name_en": "Azerbaijan",
	  "code": "AZ"
	},
	{
	  "name_pl": "Bahamy",
	  "name_en": "Bahamas",
	  "code": "BS"
	},
	{
	  "name_pl": "Bahrajn",
	  "name_en": "Bahrain",
	  "code": "BH"
	},
	{
	  "name_pl": "Bangladesz",
	  "name_en": "Bangladesh",
	  "code": "BD"
	},
	{
	  "name_pl": "Barbados",
	  "name_en": "Barbados",
	  "code": "BB"
	},
	{
	  "name_pl": "Belgia",
	  "name_en": "Belgium",
	  "code": "BE"
	},
	{
	  "name_pl": "Belize",
	  "name_en": "Belize",
	  "code": "BZ"
	},
	{
	  "name_pl": "Benin",
	  "name_en": "Benin",
	  "code": "BJ"
	},
	{
	  "name_pl": "Bermudy",
	  "name_en": "Bermuda",
	  "code": "BM"
	},
	{
	  "name_pl": "Bhutan",
	  "name_en": "Bhutan",
	  "code": "BT"
	},
	{
	  "name_pl": "Białoruś",
	  "name_en": "Belarus",
	  "code": "BY"
	},
	{
	  "name_pl": "Boliwia",
	  "name_en": "Bolivia, Plurinational State of",
	  "code": "BO"
	},
	{
	  "name_pl": "Bonaire, Sint Eustatius i Saba",
	  "name_en": "Bonaire, Sint Eustatius and Saba",
	  "code": "BQ"
	},
	{
	  "name_pl": "Bośnia i Hercegowina",
	  "name_en": "Bosnia and Herzegovina",
	  "code": "BA"
	},
	{
	  "name_pl": "Botswana",
	  "name_en": "Botswana",
	  "code": "BW"
	},
	{
	  "name_pl": "Brazylia",
	  "name_en": "Brazil",
	  "code": "BR"
	},
	{
	  "name_pl": "Brunei",
	  "name_en": "Brunei Darussalam",
	  "code": "BN"
	},
	{
	  "name_pl": "Brytyjskie Terytorium Oceanu Indyjskiego",
	  "name_en": "British Indian Ocean Territory",
	  "code": "IO"
	},
	{
	  "name_pl": "Brytyjskie Wyspy Dziewicze",
	  "name_en": "Virgin Islands, British",
	  "code": "VG"
	},
	{
	  "name_pl": "Bułgaria",
	  "name_en": "Bulgaria",
	  "code": "BG"
	},
	{
	  "name_pl": "Burkina Faso",
	  "name_en": "Burkina Faso",
	  "code": "BF"
	},
	{
	  "name_pl": "Burundi",
	  "name_en": "Burundi",
	  "code": "BI"
	},
	{
	  "name_pl": "Chile",
	  "name_en": "Chile",
	  "code": "CL"
	},
	{
	  "name_pl": "Chiny",
	  "name_en": "China",
	  "code": "CN"
	},
	{
	  "name_pl": "Chorwacja",
	  "name_en": "Croatia",
	  "code": "HR"
	},
	{
	  "name_pl": "Curaçao",
	  "name_en": "Curaçao",
	  "code": "CW"
	},
	{
	  "name_pl": "Cypr",
	  "name_en": "Cyprus",
	  "code": "CY"
	},
	{
	  "name_pl": "Czad",
	  "name_en": "Chad",
	  "code": "TD"
	},
	{
	  "name_pl": "Czarnogóra",
	  "name_en": "Montenegro",
	  "code": "ME"
	},
	{
	  "name_pl": "Czechy",
	  "name_en": "Czech Republic",
	  "code": "CZ"
	},
	{
	  "name_pl": "Dalekie Wyspy Mniejsze Stanów Zjednoczonych",
	  "name_en": "United States Minor Outlying Islands",
	  "code": "UM"
	},
	{
	  "name_pl": "Dania",
	  "name_en": "Denmark",
	  "code": "DK"
	},
	{
	  "name_pl": "Demokratyczna Republika Konga",
	  "name_en": "Congo, the Democratic Republic of the",
	  "code": "CD"
	},
	{
	  "name_pl": "Dominika",
	  "name_en": "Dominica",
	  "code": "DM"
	},
	{
	  "name_pl": "Dominikana",
	  "name_en": "Dominican Republic",
	  "code": "DO"
	},
	{
	  "name_pl": "Dżibuti",
	  "name_en": "Djibouti",
	  "code": "DJ"
	},
	{
	  "name_pl": "Egipt",
	  "name_en": "Egypt",
	  "code": "EG"
	},
	{
	  "name_pl": "Ekwador",
	  "name_en": "Ecuador",
	  "code": "EC"
	},
	{
	  "name_pl": "Erytrea",
	  "name_en": "Eritrea",
	  "code": "ER"
	},
	{
	  "name_pl": "Estonia",
	  "name_en": "Estonia",
	  "code": "EE"
	},
	{
	  "name_pl": "Etiopia",
	  "name_en": "Ethiopia",
	  "code": "ET"
	},
	{
	  "name_pl": "Falklandy",
	  "name_en": "Falkland Islands (Malvinas)",
	  "code": "FK"
	},
	{
	  "name_pl": "Fidżi",
	  "name_en": "Fiji",
	  "code": "FJ"
	},
	{
	  "name_pl": "Filipiny",
	  "name_en": "Philippines",
	  "code": "PH"
	},
	{
	  "name_pl": "Finlandia",
	  "name_en": "Finland",
	  "code": "FI"
	},
	{
	  "name_pl": "Francja",
	  "name_en": "France",
	  "code": "FR"
	},
	{
	  "name_pl": "Francuskie Terytoria Południowe i Antarktyczne",
	  "name_en": "French Southern Territories",
	  "code": "TF"
	},
	{
	  "name_pl": "Gabon",
	  "name_en": "Gabon",
	  "code": "GA"
	},
	{
	  "name_pl": "Gambia",
	  "name_en": "Gambia",
	  "code": "GM"
	},
	{
	  "name_pl": "Georgia Południowa i Sandwich Południowy",
	  "name_en": "South Georgia and the South Sandwich Islands",
	  "code": "GS"
	},
	{
	  "name_pl": "Ghana",
	  "name_en": "Ghana",
	  "code": "GH"
	},
	{
	  "name_pl": "Gibraltar",
	  "name_en": "Gibraltar",
	  "code": "GI"
	},
	{
	  "name_pl": "Grecja",
	  "name_en": "Greece",
	  "code": "GR"
	},
	{
	  "name_pl": "Grenada",
	  "name_en": "Grenada",
	  "code": "GD"
	},
	{
	  "name_pl": "Grenlandia",
	  "name_en": "Greenland",
	  "code": "GL"
	},
	{
	  "name_pl": "Gruzja",
	  "name_en": "Georgia",
	  "code": "GE"
	},
	{
	  "name_pl": "Guam",
	  "name_en": "Guam",
	  "code": "GU"
	},
	{
	  "name_pl": "Guernsey",
	  "name_en": "Guernsey",
	  "code": "GG"
	},
	{
	  "name_pl": "Gujana Francuska",
	  "name_en": "French Guiana",
	  "code": "GF"
	},
	{
	  "name_pl": "Gujana",
	  "name_en": "Guyana",
	  "code": "GY"
	},
	{
	  "name_pl": "Gwadelupa",
	  "name_en": "Guadeloupe",
	  "code": "GP"
	},
	{
	  "name_pl": "Gwatemala",
	  "name_en": "Guatemala",
	  "code": "GT"
	},
	{
	  "name_pl": "Gwinea Bissau",
	  "name_en": "Guinea-Bissau",
	  "code": "GW"
	},
	{
	  "name_pl": "Gwinea Równikowa",
	  "name_en": "Equatorial Guinea",
	  "code": "GQ"
	},
	{
	  "name_pl": "Gwinea",
	  "name_en": "Guinea",
	  "code": "GN"
	},
	{
	  "name_pl": "Haiti",
	  "name_en": "Haiti",
	  "code": "HT"
	},
	{
	  "name_pl": "Hiszpania",
	  "name_en": "Spain",
	  "code": "ES"
	},
	{
	  "name_pl": "Holandia",
	  "name_en": "Netherlands",
	  "code": "NL"
	},
	{
	  "name_pl": "Honduras",
	  "name_en": "Honduras",
	  "code": "HN"
	},
	{
	  "name_pl": "Hongkong",
	  "name_en": "Hong Kong",
	  "code": "HK"
	},
	{
	  "name_pl": "Indie",
	  "name_en": "India",
	  "code": "IN"
	},
	{
	  "name_pl": "Indonezja",
	  "name_en": "Indonesia",
	  "code": "ID"
	},
	{
	  "name_pl": "Irak",
	  "name_en": "Iraq",
	  "code": "IQ"
	},
	{
	  "name_pl": "Iran",
	  "name_en": "Iran, Islamic Republic of",
	  "code": "IR"
	},
	{
	  "name_pl": "Irlandia",
	  "name_en": "Ireland",
	  "code": "IE"
	},
	{
	  "name_pl": "Islandia",
	  "name_en": "Iceland",
	  "code": "IS"
	},
	{
	  "name_pl": "Izrael",
	  "name_en": "Israel",
	  "code": "IL"
	},
	{
	  "name_pl": "Jamajka",
	  "name_en": "Jamaica",
	  "code": "JM"
	},
	{
	  "name_pl": "Japonia",
	  "name_en": "Japan",
	  "code": "JP"
	},
	{
	  "name_pl": "Jemen",
	  "name_en": "Yemen",
	  "code": "YE"
	},
	{
	  "name_pl": "Jersey",
	  "name_en": "Jersey",
	  "code": "JE"
	},
	{
	  "name_pl": "Jordania",
	  "name_en": "Jordan",
	  "code": "JO"
	},
	{
	  "name_pl": "Kajmany",
	  "name_en": "Cayman Islands",
	  "code": "KY"
	},
	{
	  "name_pl": "Kambodża",
	  "name_en": "Cambodia",
	  "code": "KH"
	},
	{
	  "name_pl": "Kamerun",
	  "name_en": "Cameroon",
	  "code": "CM"
	},
	{
	  "name_pl": "Kanada",
	  "name_en": "Canada",
	  "code": "CA"
	},
	{
	  "name_pl": "Katar",
	  "name_en": "Qatar",
	  "code": "QA"
	},
	{
	  "name_pl": "Kazachstan",
	  "name_en": "Kazakhstan",
	  "code": "KZ"
	},
	{
	  "name_pl": "Kenia",
	  "name_en": "Kenya",
	  "code": "KE"
	},
	{
	  "name_pl": "Kirgistan",
	  "name_en": "Kyrgyzstan",
	  "code": "KG"
	},
	{
	  "name_pl": "Kiribati",
	  "name_en": "Kiribati",
	  "code": "KI"
	},
	{
	  "name_pl": "Kolumbia",
	  "name_en": "Colombia",
	  "code": "CO"
	},
	{
	  "name_pl": "Komory",
	  "name_en": "Comoros",
	  "code": "KM"
	},
	{
	  "name_pl": "Kongo",
	  "name_en": "Congo",
	  "code": "CG"
	},
	{
	  "name_pl": "Korea Południowa",
	  "name_en": "Korea, Republic of",
	  "code": "KR"
	},
	{
	  "name_pl": "Korea Północna",
	  "name_en": "Korea, Democratic People's Republic of",
	  "code": "KP"
	},
	{
	  "name_pl": "Kostaryka",
	  "name_en": "Costa Rica",
	  "code": "CR"
	},
	{
	  "name_pl": "Kuba",
	  "name_en": "Cuba",
	  "code": "CU"
	},
	{
	  "name_pl": "Kuwejt",
	  "name_en": "Kuwait",
	  "code": "KW"
	},
	{
	  "name_pl": "Laos",
	  "name_en": "Lao People's Democratic Republic",
	  "code": "LA"
	},
	{
	  "name_pl": "Lesotho",
	  "name_en": "Lesotho",
	  "code": "LS"
	},
	{
	  "name_pl": "Liban",
	  "name_en": "Lebanon",
	  "code": "LB"
	},
	{
	  "name_pl": "Liberia",
	  "name_en": "Liberia",
	  "code": "LR"
	},
	{
	  "name_pl": "Libia",
	  "name_en": "Libyan Arab Jamahiriya",
	  "code": "LY"
	},
	{
	  "name_pl": "Liechtenstein",
	  "name_en": "Liechtenstein",
	  "code": "LI"
	},
	{
	  "name_pl": "Litwa",
	  "name_en": "Lithuania",
	  "code": "LT"
	},
	{
	  "name_pl": "Luksemburg",
	  "name_en": "Luxembourg",
	  "code": "LU"
	},
	{
	  "name_pl": "Łotwa",
	  "name_en": "Latvia",
	  "code": "LV"
	},
	{
	  "name_pl": "Macedonia",
	  "name_en": "Macedonia, the former Yugoslav Republic of",
	  "code": "MK"
	},
	{
	  "name_pl": "Madagaskar",
	  "name_en": "Madagascar",
	  "code": "MG"
	},
	{
	  "name_pl": "Majotta",
	  "name_en": "Mayotte",
	  "code": "YT"
	},
	{
	  "name_pl": "Makau",
	  "name_en": "Macao",
	  "code": "MO"
	},
	{
	  "name_pl": "Malawi",
	  "name_en": "Malawi",
	  "code": "MW"
	},
	{
	  "name_pl": "Malediwy",
	  "name_en": "Maldives",
	  "code": "MV"
	},
	{
	  "name_pl": "Malezja",
	  "name_en": "Malaysia",
	  "code": "MY"
	},
	{
	  "name_pl": "Mali",
	  "name_en": "Mali",
	  "code": "ML"
	},
	{
	  "name_pl": "Malta",
	  "name_en": "Malta",
	  "code": "MT"
	},
	{
	  "name_pl": "Mariany Północne",
	  "name_en": "Northern Mariana Islands",
	  "code": "MP"
	},
	{
	  "name_pl": "Maroko",
	  "name_en": "Morocco",
	  "code": "MA"
	},
	{
	  "name_pl": "Martynika",
	  "name_en": "Martinique",
	  "code": "MQ"
	},
	{
	  "name_pl": "Mauretania",
	  "name_en": "Mauritania",
	  "code": "MR"
	},
	{
	  "name_pl": "Mauritius",
	  "name_en": "Mauritius",
	  "code": "MU"
	},
	{
	  "name_pl": "Meksyk",
	  "name_en": "Mexico",
	  "code": "MX"
	},
	{
	  "name_pl": "Mikronezja",
	  "name_en": "Micronesia, Federated States of",
	  "code": "FM"
	},
	{
	  "name_pl": "Mjanma",
	  "name_en": "Myanmar",
	  "code": "MM"
	},
	{
	  "name_pl": "Mołdawia",
	  "name_en": "Moldova, Republic of",
	  "code": "MD"
	},
	{
	  "name_pl": "Monako",
	  "name_en": "Monaco",
	  "code": "MC"
	},
	{
	  "name_pl": "Mongolia",
	  "name_en": "Mongolia",
	  "code": "MN"
	},
	{
	  "name_pl": "Montserrat",
	  "name_en": "Montserrat",
	  "code": "MS"
	},
	{
	  "name_pl": "Mozambik",
	  "name_en": "Mozambique",
	  "code": "MZ"
	},
	{
	  "name_pl": "Namibia",
	  "name_en": "Namibia",
	  "code": "NA"
	},
	{
	  "name_pl": "Nauru",
	  "name_en": "Nauru",
	  "code": "NR"
	},
	{
	  "name_pl": "Nepal",
	  "name_en": "Nepal",
	  "code": "NP"
	},
	{
	  "name_pl": "Niemcy",
	  "name_en": "Germany",
	  "code": "DE"
	},
	{
	  "name_pl": "Niger",
	  "name_en": "Niger",
	  "code": "NE"
	},
	{
	  "name_pl": "Nigeria",
	  "name_en": "Nigeria",
	  "code": "NG"
	},
	{
	  "name_pl": "Nikaragua",
	  "name_en": "Nicaragua",
	  "code": "NI"
	},
	{
	  "name_pl": "Niue",
	  "name_en": "Niue",
	  "code": "NU"
	},
	{
	  "name_pl": "Norfolk",
	  "name_en": "Norfolk Island",
	  "code": "NF"
	},
	{
	  "name_pl": "Norwegia",
	  "name_en": "Norway",
	  "code": "NO"
	},
	{
	  "name_pl": "Nowa Kaledonia",
	  "name_en": "New Caledonia",
	  "code": "NC"
	},
	{
	  "name_pl": "Nowa Zelandia",
	  "name_en": "New Zealand",
	  "code": "NZ"
	},
	{
	  "name_pl": "Oman",
	  "name_en": "Oman",
	  "code": "OM"
	},
	{
	  "name_pl": "Pakistan",
	  "name_en": "Pakistan",
	  "code": "PK"
	},
	{
	  "name_pl": "Palau",
	  "name_en": "Palau",
	  "code": "PW"
	},
	{
	  "name_pl": "Palestyna",
	  "name_en": "Palestinian Territory, Occupied",
	  "code": "PS"
	},
	{
	  "name_pl": "Panama",
	  "name_en": "Panama",
	  "code": "PA"
	},
	{
	  "name_pl": "Papua-Nowa Gwinea",
	  "name_en": "Papua New Guinea",
	  "code": "PG"
	},
	{
	  "name_pl": "Paragwaj",
	  "name_en": "Paraguay",
	  "code": "PY"
	},
	{
	  "name_pl": "Peru",
	  "name_en": "Peru",
	  "code": "PE"
	},
	{
	  "name_pl": "Pitcairn",
	  "name_en": "Pitcairn",
	  "code": "PN"
	},
	{
	  "name_pl": "Polinezja Francuska",
	  "name_en": "French Polynesia",
	  "code": "PF"
	},
	{
	  "name_pl": "Polska",
	  "name_en": "Poland",
	  "code": "PL"
	},
	{
	  "name_pl": "Portoryko",
	  "name_en": "Puerto Rico",
	  "code": "PR"
	},
	{
	  "name_pl": "Portugalia",
	  "name_en": "Portugal",
	  "code": "PT"
	},
	{
	  "name_pl": "Republika Południowej Afryki",
	  "name_en": "South Africa",
	  "code": "ZA"
	},
	{
	  "name_pl": "Republika Środkowoafrykańska",
	  "name_en": "Central African Republic",
	  "code": "CF"
	},
	{
	  "name_pl": "Republika Zielonego Przylądka",
	  "name_en": "Cape Verde",
	  "code": "CV"
	},
	{
	  "name_pl": "Reunion",
	  "name_en": "Réunion",
	  "code": "RE"
	},
	{
	  "name_pl": "Rosja",
	  "name_en": "Russian Federation",
	  "code": "RU"
	},
	{
	  "name_pl": "Rumunia",
	  "name_en": "Romania",
	  "code": "RO"
	},
	{
	  "name_pl": "Rwanda",
	  "name_en": "Rwanda",
	  "code": "RW"
	},
	{
	  "name_pl": "Sahara Zachodnia",
	  "name_en": "Western Sahara",
	  "code": "EH"
	},
	{
	  "name_pl": "Saint Kitts i Nevis",
	  "name_en": "Saint Kitts and Nevis",
	  "code": "KN"
	},
	{
	  "name_pl": "Saint Lucia",
	  "name_en": "Saint Lucia",
	  "code": "LC"
	},
	{
	  "name_pl": "Saint Vincent i Grenadyny",
	  "name_en": "Saint Vincent and the Grenadines",
	  "code": "VC"
	},
	{
	  "name_pl": "Saint-Barthélemy",
	  "name_en": "Saint Barthélemy",
	  "code": "BL"
	},
	{
	  "name_pl": "Saint-Martin",
	  "name_en": "Saint Martin (French part)",
	  "code": "MF"
	},
	{
	  "name_pl": "Saint-Pierre i Miquelon",
	  "name_en": "Saint Pierre and Miquelon",
	  "code": "PM"
	},
	{
	  "name_pl": "Salwador",
	  "name_en": "El Salvador",
	  "code": "SV"
	},
	{
	  "name_pl": "Samoa Amerykańskie",
	  "name_en": "American Samoa",
	  "code": "AS"
	},
	{
	  "name_pl": "Samoa",
	  "name_en": "Samoa",
	  "code": "WS"
	},
	{
	  "name_pl": "San Marino",
	  "name_en": "San Marino",
	  "code": "SM"
	},
	{
	  "name_pl": "Senegal",
	  "name_en": "Senegal",
	  "code": "SN"
	},
	{
	  "name_pl": "Serbia",
	  "name_en": "Serbia",
	  "code": "RS"
	},
	{
	  "name_pl": "Seszele",
	  "name_en": "Seychelles",
	  "code": "SC"
	},
	{
	  "name_pl": "Sierra Leone",
	  "name_en": "Sierra Leone",
	  "code": "SL"
	},
	{
	  "name_pl": "Singapur",
	  "name_en": "Singapore",
	  "code": "SG"
	},
	{
	  "name_pl": "Sint Maarten",
	  "name_en": "Sint Maarten (Dutch part)",
	  "code": "SX"
	},
	{
	  "name_pl": "Słowacja",
	  "name_en": "Slovakia",
	  "code": "SK"
	},
	{
	  "name_pl": "Słowenia",
	  "name_en": "Slovenia",
	  "code": "SI"
	},
	{
	  "name_pl": "Somalia",
	  "name_en": "Somalia",
	  "code": "SO"
	},
	{
	  "name_pl": "Sri Lanka",
	  "name_en": "Sri Lanka",
	  "code": "LK"
	},
	{
	  "name_pl": "Stany Zjednoczone",
	  "name_en": "United States",
	  "code": "US"
	},
	{
	  "name_pl": "Suazi",
	  "name_en": "Swaziland",
	  "code": "SZ"
	},
	{
	  "name_pl": "Sudan",
	  "name_en": "Sudan",
	  "code": "SD"
	},
	{
	  "name_pl": "Sudan Południowy",
	  "name_en": "South Sudan",
	  "code": "SS"
	},
	{
	  "name_pl": "Surinam",
	  "name_en": "Suriname",
	  "code": "SR"
	},
	{
	  "name_pl": "Svalbard i Jan Mayen",
	  "name_en": "Svalbard and Jan Mayen",
	  "code": "SJ"
	},
	{
	  "name_pl": "Syria",
	  "name_en": "Syrian Arab Republic",
	  "code": "SY"
	},
	{
	  "name_pl": "Szwajcaria",
	  "name_en": "Switzerland",
	  "code": "CH"
	},
	{
	  "name_pl": "Szwecja",
	  "name_en": "Sweden",
	  "code": "SE"
	},
	{
	  "name_pl": "Tadżykistan",
	  "name_en": "Tajikistan",
	  "code": "TJ"
	},
	{
	  "name_pl": "Tajlandia",
	  "name_en": "Thailand",
	  "code": "TH"
	},
	{
	  "name_pl": "Tajwan",
	  "name_en": "Taiwan, Province of China",
	  "code": "TW"
	},
	{
	  "name_pl": "Tanzania",
	  "name_en": "Tanzania, United Republic of",
	  "code": "TZ"
	},
	{
	  "name_pl": "Timor Wschodni",
	  "name_en": "Timor-Leste",
	  "code": "TL"
	},
	{
	  "name_pl": "Togo",
	  "name_en": "Togo",
	  "code": "TG"
	},
	{
	  "name_pl": "Tokelau",
	  "name_en": "Tokelau",
	  "code": "TK"
	},
	{
	  "name_pl": "Tonga",
	  "name_en": "Tonga",
	  "code": "TO"
	},
	{
	  "name_pl": "Trynidad i Tobago",
	  "name_en": "Trinidad and Tobago",
	  "code": "TT"
	},
	{
	  "name_pl": "Tunezja",
	  "name_en": "Tunisia",
	  "code": "TN"
	},
	{
	  "name_pl": "Turcja",
	  "name_en": "Turkey",
	  "code": "TR"
	},
	{
	  "name_pl": "Turkmenistan",
	  "name_en": "Turkmenistan",
	  "code": "TM"
	},
	{
	  "name_pl": "Turks i Caicos",
	  "name_en": "Turks and Caicos Islands",
	  "code": "TC"
	},
	{
	  "name_pl": "Tuvalu",
	  "name_en": "Tuvalu",
	  "code": "TV"
	},
	{
	  "name_pl": "Uganda",
	  "name_en": "Uganda",
	  "code": "UG"
	},
	{
	  "name_pl": "Ukraina",
	  "name_en": "Ukraine",
	  "code": "UA"
	},
	{
	  "name_pl": "Urugwaj",
	  "name_en": "Uruguay",
	  "code": "UY"
	},
	{
	  "name_pl": "Uzbekistan",
	  "name_en": "Uzbekistan",
	  "code": "UZ"
	},
	{
	  "name_pl": "Vanuatu",
	  "name_en": "Vanuatu",
	  "code": "VU"
	},
	{
	  "name_pl": "Wallis i Futuna",
	  "name_en": "Wallis and Futuna",
	  "code": "WF"
	},
	{
	  "name_pl": "Watykan",
	  "name_en": "Holy See (Vatican City State)",
	  "code": "VA"
	},
	{
	  "name_pl": "Wenezuela",
	  "name_en": "Venezuela, Bolivarian Republic of",
	  "code": "VE"
	},
	{
	  "name_pl": "Węgry",
	  "name_en": "Hungary",
	  "code": "HU"
	},
	{
	  "name_pl": "Anglia",
	  "name_en": "United Kingdom",
	  "code": "GB"
	},
	{
	  "name_pl": "Wietnam",
	  "name_en": "Viet Nam",
	  "code": "VN"
	},
	{
	  "name_pl": "Włochy",
	  "name_en": "Italy",
	  "code": "IT"
	},
	{
	  "name_pl": "Wybrzeże Kości Słoniowej",
	  "name_en": "Côte d'Ivoire",
	  "code": "CI"
	},
	{
	  "name_pl": "Wyspa Bouveta",
	  "name_en": "Bouvet Island",
	  "code": "BV"
	},
	{
	  "name_pl": "Wyspa Bożego Narodzenia",
	  "name_en": "Christmas Island",
	  "code": "CX"
	},
	{
	  "name_pl": "Wyspa Man",
	  "name_en": "Isle of Man",
	  "code": "IM"
	},
	{
	  "name_pl": "Wyspa Świętej Heleny, Wyspa Wniebowstąpienia i Tristan da Cunha",
	  "name_en": "Saint Helena, Ascension and Tristan Cunha",
	  "code": "SH"
	},
	{
	  "name_pl": "Wyspy Alandzkie",
	  "name_en": "Åland Islands",
	  "code": "AX"
	},
	{
	  "name_pl": "Wyspy Cooka",
	  "name_en": "Cook Islands",
	  "code": "CK"
	},
	{
	  "name_pl": "Wyspy Dziewicze Stanów Zjednoczonych",
	  "name_en": "Virgin Islands, U.S.",
	  "code": "VI"
	},
	{
	  "name_pl": "Wyspy Heard i McDonalda",
	  "name_en": "Heard Island and McDonald Islands",
	  "code": "HM"
	},
	{
	  "name_pl": "Wyspy Kokosowe",
	  "name_en": "Cocos (Keeling) Islands",
	  "code": "CC"
	},
	{
	  "name_pl": "Wyspy Marshalla",
	  "name_en": "Marshall Islands",
	  "code": "MH"
	},
	{
	  "name_pl": "Wyspy Owcze",
	  "name_en": "Faroe Islands",
	  "code": "FO"
	},
	{
	  "name_pl": "Wyspy Salomona",
	  "name_en": "Solomon Islands",
	  "code": "SB"
	},
	{
	  "name_pl": "Wyspy Świętego Tomasza i Książęca",
	  "name_en": "Sao Tome and Principe",
	  "code": "ST"
	},
	{
	  "name_pl": "Zambia",
	  "name_en": "Zambia",
	  "code": "ZM"
	},
	{
	  "name_pl": "Zimbabwe",
	  "name_en": "Zimbabwe",
	  "code": "ZW"
	},
	{
	  "name_pl": "Zjednoczone Emiraty Arabskie",
	  "name_en": "United Arab Emirates",
	  "code": "AE"
	}
  ]`



var countries []Country = ParseCountries()

type Country struct {
	NamePl string `json:"name_pl"`
	NameEn string `json:"name_en"`
	Code   string `json:"code"`
}

func ParseCountries() []Country {
	var data []Country
	if err := json.Unmarshal([]byte(JSON), &data); err != nil {
		panic(err)
	}

	return data
}


func MapCountryToIso(name string) string {
	for _, e := range countries {
		if e.NamePl == name {
			return e.Code
		}
	}
	fmt.Printf("Unable to find country %s\n", name)
	return name
}

func MapCodeToCountry(code string) string {
	code = strings.ToUpper(code)
	for _, e := range countries {
		if e.Code == code {
			return e.NamePl
		}
	}
	fmt.Printf("Unable to find country %s\n", code)
	return code
}