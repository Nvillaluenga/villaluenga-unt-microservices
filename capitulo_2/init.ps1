Start-Job -Scriptblock {
  cd C:\Users\nacho\Desktop\Nacho\UNT\'Proyecto final'\git\capitulo_2\cobros
  go run main.go
}
Start-Job -Scriptblock {
  cd C:\Users\nacho\Desktop\Nacho\UNT\'Proyecto final'\git\capitulo_2\stock
  .\.venv\Scripts\activate
  python .\main.py
}