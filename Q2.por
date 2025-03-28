programa {
  inclua biblioteca Matematica --> mat
  real n, x, y, z, a, b, i 
  funcao inicio() {
   
 escreva("Digite o número de testes = ")
  leia(n)
 real vetor[n]
para (i = 0; i<= n-1; i++) { escreva("Número de pessoas que compraram ingresso = ")
  leia(x)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Popular = ")
  leia(y)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Geral = ")
  leia(z)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Arquibancada = ")
  leia(a)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Cadeiras = ")
  leia(b)
  vetor[i] = (y /100) * (x* 1) + (z / 100) * (x* 5) + (a / 100) * (x * 10) + (b/ 100) * (x * 20)}
  para (i= 0; i<= n-1; i++) {escreva("A RENDA DO JOGO ", i + 1, " E = ", mat.arredondar(vetor[i],2),"\n")
  } 
}
}