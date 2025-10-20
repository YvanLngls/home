void setup() {
  // Initialise la broche 2 comme sortie
  pinMode(2, OUTPUT); 
}

void loop() {
  // Allume la LED
  digitalWrite(2, HIGH);  
  delay(1000);              // Attend une seconde
  
  // Éteint la LED
  digitalWrite(2, LOW);   
  delay(1000);              // Attend une seconde
}
