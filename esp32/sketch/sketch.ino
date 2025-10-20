#include <Adafruit_NeoPixel.h>

// Définissez la broche LED (GPIO 47 pour la LED RGB)
#define LED_PIN    47
// Définissez le nombre de LED (une seule)
#define LED_COUNT  1

// Créez l'objet NeoPixel
Adafruit_NeoPixel pixel(LED_COUNT, LED_PIN, NEO_GRB + NEO_KHZ800);

void setup() {
  // Initialise la communication série à 115200 bauds (standard)
  Serial.begin(115200); 
  Serial.println("--- Démarrage du Sketch ---");
  
  pixel.begin(); // Initialise l'objet NeoPixel
  Serial.println("LED NeoPixel initialisée sur GPIO 47.");
}

void loop() {
  // 1. Allumer la LED en ROUGE
  pixel.setPixelColor(0, pixel.Color(255, 0, 0)); 
  pixel.show(); 
  Serial.println("LED allumée (Rouge).");
  delay(1000); 

  // 2. Éteindre la LED
  pixel.setPixelColor(0, pixel.Color(0, 0, 0)); 
  pixel.show(); 
  Serial.println("LED éteinte.");
  delay(1000); 
}