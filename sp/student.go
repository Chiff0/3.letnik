import "fmt"


type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}

type Config struct {
	MinOcena int
	MaxOcena int
	StOcen   int
}

func DefaultConfig() Config {
	return Config{
		MinOcena: 6,
		MaxOcena: 10,
		StOcen:   1,
	}
}




func DodajOceno (studenti map[string]Student, vpisnaStevilka string, ocena int)
{
	if ocena < 6 || ocena > 10
	{
		fmt.print ("Narobne ocene bozo")
	}
	else 
	{
		student, ok = studenti[vpisnaStevilka]
		if (ok)
		{
			student.ocene[len[ocene]] = ocena
		}
		else 
		{
			fmt.print ("Student ne obstaja")
		}
	}
}



func IzpisRedovalnice (studenti map[string]Student)
{
	fmt.print ("REDOVALNICA:\n")

	for key, value := range studenti
	{
		fmt.Printf ("%s: %+v\n", key, value);
	}
}

func IzpisiKoncniUspeh (studenti map[string]Student)
{
	for key, value := range studenti
	{
		ocena := povprecje (studenti, key)

		fmt.Printf ("%s: povprečna ocena %d -> ", value.ime, ocena)
		
		if ocena >= 9
		{
			fmt.Printf ("ODLIČEN ŠTUDENT!!!!!\n")
		}
		else if ocena >= 6
		{
			fmt.Printf ("Povprečen študent :|\n")
		}
		else
		{
			fmt.Printf ("SLAB študent >: \n")
		}
	}
}






