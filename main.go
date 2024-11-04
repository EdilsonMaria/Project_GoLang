//Servidor HTTP simples escrito em Go
package main

//Importando os packager nescessarios (onverter os dados em Json, configurar o servidor HTTP e formatar datas e horários.)
import (
    "encoding/json"
    "net/http"
    "time"
)

//Response para representar os dados que vamos enviar como resposta ao cliente.
type Response struct {
    Nome    string `json:"nome"`
    Horario string `json:"horario"`
}

//Handler é chamada sempre que uma requisição é feita ao servidor.
func handler(w http.ResponseWriter, r *http.Request) {
    location, err := time.LoadLocation("America/Sao_Paulo")
    if err != nil {
        http.Error(w, "Erro ao carregar o fuso horário", http.StatusInternalServerError)
        return
    }

    //Define o horário de Brasília
    response := Response{
        Nome:    "Projeto Korp",
        Horario: time.Now().In(location).Format(time.RFC3339),
    }
    
	//Configurar e enviar a resposta
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}