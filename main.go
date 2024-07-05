package main

import (
    "fmt"
    "strconv"
    // "unsafe"
)

func main() {

    {// VARIABLES ---------------------


    // var myInt int
    // myInt = 12
    // fmt.Println("mi entero es: ", myInt)

    // i := 1032049348
    // fmt.Println(i)

    // myString := "este es mi string"
    // fmt.Println(myString) 

	// fmt.Println("Fin del programa")

    // const myConst int = 21
    // fmt.Println("esta es mi constante: ", myConst)

    // const myFirtsStringConst = "una constante que no se puede renombrar y si no se le indica el tipo de dato lo coloca auto"
    // fmt.Println(myFirtsStringConst)
    }


    {// SCOPES -----------------------

    // myOtherScopeVariable := 50
    // {
    //     myScopeVariable := 40
    //     {
    //         fmt.Println(myOtherScopeVariable)
    //         fmt.Println(myScopeVariable)
    //     }
    // }
    // fmt.Println(myOtherScopeVariable)
    }


    {// UNIT VARIABLE SIZE ------------------------

    // var my8BitsUintvar uint8 = 20
    // fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUintvar, my8BitsUintvar, unsafe.Sizeof(my8BitsUintvar), unsafe.Sizeof(my8BitsUintvar)*8)

    // var my16BitsUintvar uint16 = 25
    // fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsUintvar, my16BitsUintvar, unsafe.Sizeof(my16BitsUintvar), unsafe.Sizeof(my16BitsUintvar)*8)

    // var my32BitsUintvar uint32 = 30
    // fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsUintvar, my32BitsUintvar, unsafe.Sizeof(my32BitsUintvar), unsafe.Sizeof(my32BitsUintvar)*8)

    // var my64BitsUintvar uint64 = 90
    // fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsUintvar, my64BitsUintvar, unsafe.Sizeof(my64BitsUintvar), unsafe.Sizeof(my64BitsUintvar)*8)
    }


    {// FLOAT VARIABLE -----------------
    
    // var myFloat32Var float32 = 32.14
    // fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

    // var myFloat64Var float64 = 590.1234
    // fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

    // myOtherFloat := 21.12
    // fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)
    }



    {
    // STRING VARIABLE ---------------

    // var myStringVar3 string = "test1234"
    // fmt.Printf("mi valor %s \n", myStringVar3)

    // myStringVar4 := `Mi valor
    // es 
    // este
    // `
    // fmt.Printf("mi valor %s \n", myStringVar4)
    }
    
    {
        // CONVERT TO STRING -----
        // floatVar := 33.11
        // fmt.Printf("type %T, value: %f \n", floatVar, floatVar)
        
        // floatStrvar := fmt.Sprintf("%f", floatVar)
        // fmt.Printf("type %T, value: %s \n", floatStrvar, floatStrvar)
    }

    {
        strIntVar2, err := strconv.Atoi("15s")
        fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar2, strIntVar2, err)

        strIntVar3, err := strconv.ParseInt("10", 10, 64)
        fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar3, strIntVar3, err)

        strIntVar4, err := strconv.ParseFloat("10.11", 64)
        fmt.Printf("type: %T, value: %f, err: %v \n", strIntVar4, strIntVar4, err)

    }
}