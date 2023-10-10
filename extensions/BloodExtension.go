package extensions

type BloodExtension interface {
	BloodType() string
	BloodRh() string
	BloodGroup() string
}
