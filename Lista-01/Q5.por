programa {
  funcao inicio() {
    inteiro conta 
    real a,b,c 
    caracter t
    escreva ( "Digite o número da conta ")
    leia (b)
    escreva ("Digite o valor de água consumida:")
    leia (a)
    escreva ( "Escreva o tipo da conta: ")
    leia (t)
    escreva ("CONTA:",b,"\n")
    se  (t=="R")
   { escreva ("VALOR DA CONTA =",5+0.05*a)}
   se (t=="C")
   {escreva ("VALOR DA CONTA =",500+(a-80)*0.25)}
   se (t=="I")
   {escreva ("VALOR DA CONTA =", 800+(a-100)*0.04)}
  }
}
