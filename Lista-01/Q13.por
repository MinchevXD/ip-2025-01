programa {
  funcao inicio() {
    real x
    escreva ("DIGITE A NOTA DO ALUNO:")
    leia (x)
    se (6>x)
    {escreva ("NOTA =", x, "  CONCEITO = D")}
    se (7.5>=x e x>6)
    {escreva ("NOTA =", x, "  CONCEITO = C")}
    se (9>=x e x>7.5)
    {escreva ("NOTA =", x, "  CONCEITO = B")}
    se  (x>9)
    {escreva ("NOTA =", x, "  CONCEITO = A")}
  }
}
