
const int limiarPressionado = 1000; // limiar entre pressionado ou não (min 0 e max 1023)

void setup() {
  Serial.begin(9600); //iniciar o monitor serial com taxa de trasmissão de 9600
}

void loop() {
  //se o valor da porta analogica atingir o limiar vai mandar a letra para o monitor serial
  if(analogRead(A0) <= limiarPressionado){
    Serial.println('A');
  }
  if(analogRead(A1) <= limiarPressionado){
    Serial.println('B');
  }
  if(analogRead(A2) <= limiarPressionado){
    Serial.println('C');
  }
  if(analogRead(A3) <= limiarPressionado){
    Serial.println('D');
  }
  if(analogRead(A4) <= limiarPressionado){
    Serial.println('E');
  }
  if(analogRead(A5) <= limiarPressionado){
    Serial.println('F');
  }
  
  delay(300);
}
