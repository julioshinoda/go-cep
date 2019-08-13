package main

import (
	"github.com/julioshinoda/polis/cep"

	"fmt"
)

func main() {
	//	payload := "relaxation=08717260&tipoCEP=ALL"

	fmt.Println(cep.GetAddressByZipcode("08717-260"))
	//	c.Visit("http://www.buscacep.correios.com.br/sistemas/buscacep/resultadoBuscaCepEndereco.cfm")

}
