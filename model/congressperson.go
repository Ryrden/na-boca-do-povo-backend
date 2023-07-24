package model

type CongressPerson struct {
    Id uint64 `json:"id" binding:"required"`
    Uri string  `json:"uri" binding:"required"`
    Nome string `json:"nome" binding:"required"`
    SiglaPartido string `json:"siglaPartido" binding:"required"`
    UriPartido string `json:"uriPartido" binding:"required"` 
    SiglaUf string `json:"siglaUf" binding:"required"`
    IdLegislatura uint64 `json:"idLegislatura" binding:"required"`
    UrlFoto string `json:"urlFoto" binding:"required"`
    Email string `json:"email" binding:"required, email"`
}