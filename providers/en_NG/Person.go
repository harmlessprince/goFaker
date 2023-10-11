package en_NG

import (
	"github.com/harmlessprince/goFaker/providers"
)

type EnNGPerson struct {
	providers.BasePerson
}

func NewPerson() *EnNGPerson {
	p := &EnNGPerson{}
	p.SetMaleNameFormats()
	p.SetFemaleNameFormats()
	p.SetFirstNameFormats()
	p.SetMaleFirstNames()
	p.SetFemaleFirstNames()
	p.SetLastNames()
	return p
}

// SetMaleNameFormats is a method that provides access to the maleNameFormats.
func (p *EnNGPerson) SetMaleNameFormats() {
	maleNameFormats := [][]func(*providers.BasePerson) string{
		{(*providers.BasePerson).FirstNameMale, (*providers.BasePerson).LastName},
		{(*providers.BasePerson).FirstNameMale, (*providers.BasePerson).FirstNameMale, (*providers.BasePerson).LastName},
	}
	p.BasePerson.SetMaleNameFormats(maleNameFormats)
}

// SetFemaleNameFormats is a method that provides access to the femaleNameFormats.
func (p *EnNGPerson) SetFemaleNameFormats() {
	femaleNameFormats := [][]func(*providers.BasePerson) string{
		{(*providers.BasePerson).FirstNameFemale, (*providers.BasePerson).LastName},
		{(*providers.BasePerson).FirstNameFemale, (*providers.BasePerson).FirstNameFemale, (*providers.BasePerson).LastName},
	}
	p.BasePerson.SetFemaleNameFormats(femaleNameFormats)
}

func (p *EnNGPerson) SetMaleFirstNames() {
	maleFirstNames := []string{
		"Abimbola", "Abisola", "Abisoye", "Adeboye", "Adedayo", "Adegoke", "Akande", "Akanni", "Alade", "Ayinde", "Azubuike",
		"Banji", "Bankole", "Buchi", "Bukola",
		"Chinedu", "Chisom", "Chukwu",
		"Damilare", "Damilola", "Danjuma",
		"Ebiowei", "Emeka", "Emmanuel", "Esse",
		"Funmilade", "Funmilayo",
		"Gbeminiyi", "Gbemisola",
		"Habiba", "Ifeanyichukwu",
		"Ikenna", "Ikhidie", "Ireti",
		"Jadesola", "Johnson",
		"Kayode", "Kemi", "Kubra", "Kubura",
		"Lolade",
		"Makinwa", "Mohammed", "Musa", "Muyiwa",
		"Nnamdi",
		"Olaide", "Olufunmi", "Olumide", "Oluwunmi", "Onoriode",
		"Remilekun", "Rotimi",
		"Shade", "Shalewa", "Sname",
		"Tari", "Temitope", "Titilope", "Tobiloba", "Toke", "Tomiloba", "Tope",
		"Uzodimma",
		"Wale",
		"Yakubu", "Yusuf", "Yusuf",
	}
	p.BasePerson.SetMaleFirstNames(maleFirstNames)
}

func (p *EnNGPerson) SetFemaleFirstNames() {
	femaleFirstNames := []string{
		"Adaugo", "Akunna", "Aminat", "Aminu", "Augustina", "Ayebatari",
		"Cherechi", "Chiamaka", "Chimamanda", "Chinyere", "Chizoba",
		"Ebiere", "Efe",
		"Fatima", "Ifeoma",
		"Ifunanya", "Isioma",
		"Jolayemi",
		"Lola",
		"Obioma", "Omawunmi", "Omolara", "Onome",
		"Rasheedah",
		"Sekinat", "Simisola", "Sumayyah",
		"Titi", "Titilayo", "Toluwani",
		"Zainab", "Adaugo", "Akunna", "Aminat", "Aminu", "Augustina", "Ayebatari",
		"Cherechi", "Chiamaka", "Chimamanda", "Chinyere", "Chizoba",
		"Ebiere", "Efe",
		"Fatima", "Ifeoma",
		"Ifunanya", "Isioma",
		"Jolayemi",
		"Lola",
		"Obioma", "Omawunmi", "Omolara", "Onome",
		"Rasheedah",
		"Sekinat", "Simisola", "Sumayyah",
		"Titi", "Titilayo", "Toluwani",
		"Zainab",
	}
	p.BasePerson.SetFemaleFirstNames(femaleFirstNames)
}

func (p *EnNGPerson) SetLastNames() {
	lastNames := []string{
		"Abiodun", "Abiola", "Abodunrin", "Abosede", "Adaobi", "Adebayo", "Adegboye", "Adegoke", "Ademayowa", "Ademola", "Adeniyan", "Adeoluwa", "Aderinsola", "Aderonke", "Adesina", "Adewale", "Adewale", "Adewale", "Adewunmi", "Adewura", "Adeyemo", "Afolabi", "Afunku", "Agboola", "Agboola", "Agnes", "Aigbiniode", "Ajakaiye", "Ajose-adeogun", "Akeem-omosanya", "Akerele", "Akintade", "Aligbe", "Amaechi", "Aminat", "Aremu", "Atanda", "Ayisat", "Ayobami", "Ayomide", "Ayomide",
		"Babalola", "Babatunde", "Balogun", "Bamisebi", "Bello", "Busari",
		"Chibike", "Chibuike", "Chidinma", "Chidozie", "Christian", "Clare",
		"David", "David",
		"Ebubechukwu", "Egbochukwu", "Ehigiator", "Ekwueme", "Elebiyo", "Elizabeth", "Elizabeth", "Elizabeth", "Emmanuel", "Emmanuel", "Esther",
		"Funmilayo",
		"Gbadamosi", "Gbogboade", "Grace",
		"Habeeb", "Hanifat", "Isaac",
		"Ismail", "Isokun", "Israel", "Iyalla",
		"Jamiu", "Jimoh", "Joshua", "Justina",
		"Katherine", "Kayode", "Kayode", "Kimberly",
		"Ladega", "Latifat", "Lawal", "Leonard",
		"Makuachukwu", "Maryam", "Maryjane", "Mayowa", "Miracle", "Mobolaji", "Mogbadunade", "Motalo", "Muinat", "Mukaram", "Mustapha", "Mutiat",
		"Ndukwu", "Ngozi", "Nojeem", "Nwachukwu", "Nwogu", "Nwuzor",
		"Obiageli", "Obianuju", "Odunayo", "Odunayo", "Ogunbanwo", "Ogunwande", "Okonkwo", "Okunola", "Oladeji", "Oladimeji", "Olaoluwa", "Olasunkanmi", "Olasunkanmi-fasayo", "Olawale", "Olubukola", "Olubunmi", "Olufeyikemi", "Olumide", "Olutola", "Oluwakemi", "Oluwanisola", "Oluwaseun", "Oluwaseyi", "Oluwashina", "Oluwatosin", "Omobolaji", "Omobolanle", "Omolara", "Omowale", "Onohinosen", "Onose", "Onyinyechukwu", "Opeyemi", "Osuagwu", "Oyebola", "Oyelude", "Oyinkansola",
		"Peter",
		"Sabdat", "Saheed", "Salami", "Samuel", "Sanusi", "Sarah", "Segunmaru", "Sekinat", "Sulaimon", "Sylvester",
		"Taiwo", "Tamunoemi", "Tella", "Temitope", "Tolulope",
		"Uchechi",
		"Wasiu", "Wilcox", "Wuraola",
		"Yaqub", "Yussuf",
	}
	p.BasePerson.SetLastNames(lastNames)
}
