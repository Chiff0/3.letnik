package main

import (
	"context"
	"fmt"
	"log"
	"os"

	student "github.com/yourusername/student-module"
	"github.com/urfave/cli/v3"
)

func main() {
	studenti := make(map[string]*student.Student)
	studenti["63220059"] = &student.Student{
		Ime:     "Filip",
		Priimek: "Dobnikar",
		Ocene:   []int{10, 10, 10, 10},
	}
	studenti["65454"] = &student.Student{
		Ime:     "Rok",
		Priimek: "Dobnikar",
		Ocene:   []int{7, 1, 2, 6},
	}
	studenti["1111111"] = &student.Student{
		Ime:     "Burek",
		Priimek: "Cmurek",
		Ocene:   []int{9, 8, 9, 9},
	}

	cmd := &cli.Command{
		Name:  "student-app",
		Usage: "Upravljanje študentov in ocen",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Value:   1,
				Usage:   "Najmanjše število ocen potrebnih za pozitivno oceno",
				Aliases: []string{"n"},
			},
			&cli.IntFlag{
				Name:    "minOcena",
				Value:   6,
				Usage:   "Najmanjša možna ocena",
				Aliases: []string{"min"},
			},
			&cli.IntFlag{
				Name:    "maxOcena",
				Value:   10,
				Usage:   "Največja možna ocena",
				Aliases: []string{"max"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "izpis",
				Usage: "Izpiši vse študente in njihove ocene",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					student.IzpisVsehOcen(studenti)
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiši končni uspeh vseh študentov",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					config := student.Config{
						MinOcena: cmd.Int("minOcena"),
						MaxOcena: cmd.Int("maxOcena"),
						StOcen:   cmd.Int("stOcen"),
					}
					fmt.Printf("Konfiguracija: minOcena=%d, maxOcena=%d, stOcen=%d\n\n", 
						config.MinOcena, config.MaxOcena, config.StOcen)
					student.IzpisiKoncniUspeh(studenti, config)
					return nil
				},
			},
			{
				Name:      "dodaj",
				Usage:     "Dodaj oceno študentu",
				ArgsUsage: "<vpisna_stevilka> <ocena>",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() != 2 {
						return fmt.Errorf("potrebujete 2 argumenta: <vpisna_stevilka> <ocena>")
					}
					
					vpisnaStevilka := cmd.Args().Get(0)
					var ocena int
					_, err := fmt.Sscanf(cmd.Args().Get(1), "%d", &ocena)
					if err != nil {
						return fmt.Errorf("napaka pri branju ocene: %v", err)
					}
					
					config := student.Config{
						MinOcena: cmd.Int("minOcena"),
						MaxOcena: cmd.Int("maxOcena"),
						StOcen:   cmd.Int("stOcen"),
					}
					
					student.DodajOceno(studenti, vpisnaStevilka, ocena, config)
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
				config := student.Config{
				MinOcena: cmd.Int("minOcena"),
				MaxOcena: cmd.Int("maxOcena"),
				StOcen:   cmd.Int("stOcen"),
			}
			fmt.Printf("Konfiguracija: minOcena=%d, maxOcena=%d, stOcen=%d\n", 
				config.MinOcena, config.MaxOcena, config.StOcen)
			student.IzpisVsehOcen(studenti)
			student.IzpisiKoncniUspeh(studenti, config)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
