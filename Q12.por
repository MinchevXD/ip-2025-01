programa {
  funcao inicio() {
    inteiro x
    escreva ("DIGITE O VALOR DE HORAS GASTAS:")
    leia (x)
    se (x%3==0)
    escreva ("O VALOR A PAGAR E =",x*10/3)
    senao
    escreva ("O VALOR A PAGAR E =",(x-x%3)*10/3 + (x%3)*5)
  }
}
