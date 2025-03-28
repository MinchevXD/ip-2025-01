programa {
  funcao inicio() {
     inteiro x, y, n, k
    real r, soma
    soma = 0
    escreva("DIGITE O PRIMEIRO TERMO, A RAZÃO E O NÚMERO DE ELEMENTOS RESPECTIVAMENTE: ")
    leia(x, y, n)
    para (k=0;k<n;k++)
    {r = x + k * y
    soma += r}
    escreva("O resultado da PA é: ", soma)
  }
}
