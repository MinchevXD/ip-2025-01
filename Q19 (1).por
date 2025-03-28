programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real n,i,soma,a
    soma=0
    escreva ("DIGITE UM NÚMERO:")
    leia (n)
    para (i=1;i<=n;i++)
    {a=1/i
    soma+=a}
    escreva (mat.arredondar(soma,2))
  }
}

