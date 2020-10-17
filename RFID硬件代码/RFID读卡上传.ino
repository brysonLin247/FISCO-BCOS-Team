#include <SPI.h>
#include <MFRC522.h>
#include <ESP8266WiFi.h>
#define RST_PIN         5           // 配置针脚
#define SS_PIN          4
#define BB_PIN          D3 
#define LED 2 //定义开发板上LED的IO
WiFiClient client;//tcp连接
MFRC522 mfrc522(SS_PIN, RST_PIN);   // 创建新的RFID实例
MFRC522::MIFARE_Key key;
int pos=0; 
int value=0;
char addr[10]="0755-PS-R";
const char* host="39.97.3.219";//服务器公网ip
WiFiServer server(80);//开启80端口

void setup() {
  Serial.begin(9600); // 设置串口波特率为9600
  SPI.begin();        // SPI开始
  mfrc522.PCD_Init(); // Init MFRC522 card
  for (byte i = 0; i < 6; i++) {
        key.keyByte[i] = 0xFF;
    }
  dump_byte_array(key.keyByte, MFRC522::MF_KEY_SIZE);
  Serial.println("test-demo-start");
  pinMode(BB_PIN,OUTPUT);
  pinMode(LED, OUTPUT);   //设置LED的工作模式
  digitalWrite(LED, 0);
  if (!AutoConfig())
  {
    SmartConfig();
  }                               //调用SmartConfig函数进行配网
  
  Serial.println("");
  Serial.println("WiFi connected");
  
  // Start the server
  server.begin();//开启服务器
  Serial.println("Server started");

  // Print the IP address
  Serial.println(WiFi.localIP());//输出板子的ip

  Serial.print("Connecting to ");
  Serial.println(host);
 
  const int httpPort=1234;//端口号
  if(!client.connect(host,httpPort)){//连接失败
    Serial.println("connection failed");
    Serial.println(WiFi.localIP());
    return;
  }
}


void loop() {
  
  // 寻找新卡
  if ( ! mfrc522.PICC_IsNewCardPresent()) {
    //Serial.println("没有找到卡");
    return;
  }

  // 选择一张卡
  if ( ! mfrc522.PICC_ReadCardSerial()) {
    Serial.println("没有卡可选");
    return;
  }

 Serial.print(F("卡片 UID:"));
  dump_byte_array(mfrc522.uid.uidByte, mfrc522.uid.size);
  Serial.println();
Serial.print(F("站点地址:"));
Serial.print(addr);



  // 显示卡片的详细信息
 
  client.print("------------------------");
  client.print(F("卡片 UID:"));
  for (byte i = 0; i < mfrc522.uid.size; i++) {
    client.print(mfrc522.uid.uidByte[i] < 0x10 ? " 0" : " ");
    client.print(mfrc522.uid.uidByte[i], HEX);
  }
client.print(F("站点地址:"));
client.print(addr);
   
  Serial.println();

 byte sector         = 1;
    byte blockAddr      = 4;
    byte dataBlock[]    = {
        0x01, 0x02, 0x03, 0x04, //  1,  2,   3,  4,
        0x05, 0x06, 0x07, 0x08, //  5,  6,   7,  8,
        0x00, 0x00, 0x00, 0x00, //  0，0，0，0
        0x00, 0x00, 0x00, 0x00  // 0，0，0，0
    };//写入的数据定义
    byte trailerBlock   = 7;
    MFRC522::StatusCode status;
    byte buffer[18];
    byte size = sizeof(buffer);

    Serial.println(F("显示所有扇区的数据"));
   mfrc522.PICC_DumpMifareClassicSectorToSerial(&(mfrc522.uid), &key, sector);
    Serial.println();
  
  digitalWrite(BB_PIN,HIGH);
  delay(100);
  digitalWrite(BB_PIN,LOW);
  delay(50);
  digitalWrite(BB_PIN,HIGH);
  delay(100);
  digitalWrite(BB_PIN,LOW);

  if (status != MFRC522::STATUS_OK) {
    Serial.print(F("身份验证失败？或者是卡链接失败"));
    Serial.println(mfrc522.GetStatusCodeName(status));
    return;
  }
  //停止 PICC
  mfrc522.PICC_HaltA();
  //停止加密PCD
  mfrc522.PCD_StopCrypto1();
  return;
}

/**
   将字节数组转储为串行的十六进制值
*/
void dump_byte_array(byte *buffer, byte bufferSize) {
  for (byte i = 0; i < bufferSize; i++) {
    Serial.print(buffer[i] < 0x10 ? " 0" : " ");
    Serial.print(buffer[i], HEX);
  }
}
void SmartConfig()
{
  WiFi.mode(WIFI_STA);
  Serial.println("\r\nWait for Smartconfig...");
  WiFi.beginSmartConfig();
  while (1)
  {
    Serial.print(".");
    delay(500);                   // wait for a second
    if (WiFi.smartConfigDone())
    {
      Serial.println("SmartConfig Success");
      Serial.printf("SSID:%s\r\n", WiFi.SSID().c_str());
      Serial.printf("PSW:%s\r\n", WiFi.psk().c_str());
      break;
    }
  }
}
bool AutoConfig()
{
    WiFi.begin();
    //如果觉得时间太长可改
    for (int i = 0; i < 20; i++)
    {
        int wstatus = WiFi.status();
        if (wstatus == WL_CONNECTED)
        {
          Serial.println("WIFI SmartConfig Success");
          Serial.printf("SSID:%s", WiFi.SSID().c_str());
          Serial.printf(", PSW:%s\r\n", WiFi.psk().c_str());
          Serial.print("LocalIP:");
          Serial.print(WiFi.localIP());
          Serial.print(" ,GateIP:");
          Serial.println(WiFi.gatewayIP());
          return true;
        }
        else
        {
          Serial.print("WIFI AutoConfig Waiting......");
          Serial.println(wstatus);
          delay(1000);
        }
    }
    Serial.println("WIFI AutoConfig Faild!" );
    return false;
}
