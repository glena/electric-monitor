#include "stdlib.h"
#include "EmonLib.h"
#include "ESP8266WiFi.h"
#include "string.h"

float networkV = 220.0;

const char* ssid = "cuchicu";
const char* password = "chuchuchuchuchu";

const char* host = "192.168.0.223";
const int   port = 6969;    

EnergyMonitor energyMonitor; 

void setup()
{
  Serial.begin(9600);
  Serial.println("");
  Serial.println("Starting");
  
  energyMonitor.current(0, 9);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  
  while (WiFi.status() != WL_CONNECTED) 
  {
     delay(500);
     Serial.print("*");
  }
  
  Serial.println("");
  Serial.println("WiFi connection Successful");
  Serial.print("The IP Address of ESP8266 Module is: ");
  Serial.println(WiFi.localIP());// Print the IP address
   
  Serial.println("Ready");
}

void loop()
{
  WiFiClient client;

  if (client.connect(host, port)) {
    Serial.println("Connected");

    double i = energyMonitor.calcIrms(1484);
    double p = i * networkV;
    
    char packet[50];
    sprintf(packet, "%f,%f;", p, i); 
    Serial.print("Message = ");
    Serial.println(packet);
    client.print(packet);

    unsigned long timeout = millis();
    while (client.available() == 0) {
      if (millis() - timeout > 5000) {
        Serial.println(">>> Client Timeout !");
        client.stop();
        delay(60000);
        return;
      }
    }

    Serial.println("receiving from remote server");
    // not testing 'client.connected()' since we do not need to send data here
    while (client.available()) {
      char ch = static_cast<char>(client.read());
      Serial.print(ch);
    }

    Serial.println("closing connection");
    client.stop();  
  }
}
