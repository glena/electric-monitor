#include "stdlib.h"
#include "EmonLib.h"
#include "string.h"

#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

float networkV = 220.0;

const char* ssid = "...";
const char* password = "...";

const char* host = "...";
const int   port = 80;    

EnergyMonitor energyMonitor; 

void setup()
{
  Serial.begin(9600);
  Serial.println("");
  Serial.println("Starting");
  
  energyMonitor.current(0, 9);

 // WiFi.mode(WIFI_STA);
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
  double i = energyMonitor.calcIrms(1484);
  double p = i * networkV;


  HTTPClient http;
  http.begin(host);
  http.addHeader("Content-Type", "application/x-www-form-urlencoded");
  http.addHeader("Authorization", "Basic ...");

  String body = String("p=" + String(p) + "&v=" + String(networkV) + "&i=" + String(i));

  Serial.print("HTTP Request body: ");
  Serial.println(body);
  
  int httpResponseCode = http.POST(body);  

  Serial.print("HTTP Response code: ");
  Serial.println(httpResponseCode);
    
  // Free resources
  http.end();

}
