programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real f,p,c,m
    escreva ("Digite o valor em fahrenheit:")
    leia (f)
    escreva ( "Digite o valor em polegadas:")
    leia (p)
    c= (5 * f - 160) / 9
    m= p * 25.4
    escreva("O VALOR EM CELSIUS = ",mat.arredondar(c,2), "\n")
    escreva("A QUANTIDADE DE CHUVA E = ",mat.arredondar(m,2) , "\n")
  }
}
