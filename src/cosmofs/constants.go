/**

Copyright (C) 2012  Roberto Costumero Moreno <roberto@costumero.es>

This file is part of Cosmofs.

Cosmofs is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Cosmofs is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Cosmofs.  If not, see <http://www.gnu.org/licenses/>.

**/

package cosmofs

import (
	"flag"
	"os"
)

const (
	COSMOFSDIR string = ".cosmofs"
	COSMOFSCONFIGFILE string = ".cosmofsconfig"
)

var (
	//Cosmofsin string = os.Getenv("COSMOFSIN")
	//Cosmofsout string = os.Getenv("COSMOFSOUT")
	Cosmofsin *string = flag.String("cosmofsin", os.Getenv("COSMOFSIN"), "Location of incoming packages")
	Cosmofsout *string = flag.String("cosmofsout", os.Getenv("COSMOFSOUT"), "Location of shared directories")
	resetConfig *bool = flag.Bool("r", false, "Re-generate config files")

	//TODO: Change prueba.pub to id_rsa.pub
	pubkeyFileName *string = flag.String("cosmofspubkey", os.Getenv("COSMOFSPUBKEY"), "Location of public RSA Key")
	privkeyFileName *string = flag.String("cosmofsprivkey", os.Getenv("COSMOFSPRIVKEY"), "Location of private RSA Key")
)

