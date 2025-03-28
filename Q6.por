programa {
  inclua biblioteca Matematica --> mat
  inteiro c
  real f, n
  funcao inicio() {
    escreva("Digite o número de conversões a serem feitas: ")
    leia(n)
    real vetor[n]
    para(c = 0; c < n; c = c +1)
   {  escreva("Digite a temperatura em fahrenheit: ")
   leia(f) 
   vetor[c] = (5/9) * ( f - 32 )}
    para(c = 0; c < n; c = c + 1) {
      escreva(vetor[c]*9/5+32,"  FAHRENHEIT EQUIVALE A  ", mat.arredondar(vetor[c],2)," CELSIUS","\n")
  }
  }
}
