import json
import os
import time
# from pymongo.mongo_client import MongoClient
# from pymongo.server_api import ServerApi

powerCombination = {}
aeroCombination = {}
lightweightCombination = {}
gripCombination = {}
toolsCombination = {}

superCombination = {}

def main():
    timeNow = time.time()
    with open("data.json") as f:
        data = json.load(f)

    
    # uri = "mongodb+srv://rashiraffi:pwd@digid.1mdserk.mongodb.net/?retryWrites=true&w=majority"
    # # Create a new client and connect to the server
    # client = MongoClient(uri, server_api=ServerApi('1'))
    # # Send a ping to confirm a successful connection
    # try:
    #     client.admin.command('ping')
    #     print("Pinged your deployment. You successfully connected to MongoDB!")
    # except Exception as e:
    #     print(e)

    # db = client["F1"]
    # collection = db["Combinations"]


    combinations = []

    power = data["power"]
    aero = data["aero"]
    lightweight = data["lightweight"]
    grip = data["grip"]
    tools = data["tools"]
    cons = data["cons"]
    maxPoints = 0
    maxConfigSet = 0
    maxConfig = 0
    
    maxIdealPoints = 0
    maxIdealConfig = 0
    maxIdealConfigSet = 0

    combiCount = 0

    for i in range(len(power)-1):
        for j in range(i+1,len(power)):
            powerCombination[str(i+1)+"-"+str(j+1)] = {
            "id": str(i+1)+"-"+str(j+1),
            "cost": power[i]["cost"] + power[j]["cost"],
            "power": power[i]["power"] + power[j]["power"],
            "aero": power[i]["aero"] + power[j]["aero"],
            "lightweight": power[i]["lightweight"] + power[j]["lightweight"],
            "grip": power[i]["grip"] + power[j]["grip"],
        }
    
    for i in range(len(aero)-1):
        for j in range(i+1,len(aero)):
            aeroCombination[str(i+1)+"-"+str(j+1)] = {
            "id": str(i+1)+"-"+str(j+1),
            "cost": aero[i]["cost"] + aero[j]["cost"],
            "power": aero[i]["power"] + aero[j]["power"],
            "aero": aero[i]["aero"] + aero[j]["aero"],
            "lightweight": aero[i]["lightweight"] + aero[j]["lightweight"],
            "grip": aero[i]["grip"] + aero[j]["grip"],
        }
    
    
    for i in range(len(lightweight)-1):
        for j in range(i+1,len(lightweight)):
            lightweightCombination[str(i+1)+"-"+str(j+1)] = {
            "id": str(i+1)+"-"+str(j+1),
            "cost": lightweight[i]["cost"] + lightweight[j]["cost"],
            "power": lightweight[i]["power"] + lightweight[j]["power"],
            "aero": lightweight[i]["aero"] + lightweight[j]["aero"],
            "lightweight": lightweight[i]["lightweight"] + lightweight[j]["lightweight"],
            "grip": lightweight[i]["grip"] + lightweight[j]["grip"],
        }
    
    
    for i in range(len(grip)-1):
        for j in range(i+1,len(grip)):
            gripCombination[str(i+1)+"-"+str(j+1)] = {
            "id": str(i+1)+"-"+str(j+1),
            "cost": grip[i]["cost"] + grip[j]["cost"],
            "power": grip[i]["power"] + grip[j]["power"],
            "aero": grip[i]["aero"] + grip[j]["aero"],
            "lightweight": grip[i]["lightweight"] + grip[j]["lightweight"],
            "grip": grip[i]["grip"] + grip[j]["grip"],
        }
    
    
    for i in range(len(tools)-1):
        for j in range(i+1,len(tools)):
            toolsCombination[str(i+1)+"-"+str(j+1)] = {
            "id": str(i+1)+"-"+str(j+1),
            "cost": tools[i]["cost"] + tools[j]["cost"],
            "power": tools[i]["power"] + tools[j]["power"],
            "aero": tools[i]["aero"] + tools[j]["aero"],
            "lightweight": tools[i]["lightweight"] + tools[j]["lightweight"],
            "grip": tools[i]["grip"] + tools[j]["grip"],
        }
            
    print(len(powerCombination)* len(aeroCombination)* len(lightweightCombination)* len(gripCombination)* len(toolsCombination)*len(cons))
    
    for pCom in powerCombination:
        for aCom in aeroCombination:
            for lCom in lightweightCombination:
                for gCom in gripCombination:
                    for tCom in toolsCombination:
                        for c in cons:
                                combiCount += 1
                                config = {
                                "cost": powerCombination[pCom]["cost"] + aeroCombination[aCom]["cost"] + lightweightCombination[lCom]["cost"] + gripCombination[gCom]["cost"] + toolsCombination[tCom]["cost"] + c["cost"],
                                "power": (107+powerCombination[pCom]["power"] + aeroCombination[aCom]["power"] + lightweightCombination[lCom]["power"] + gripCombination[gCom]["power"] + toolsCombination[tCom]["power"])*c["power"],
                                "aero": (161+powerCombination[pCom]["aero"] + aeroCombination[aCom]["aero"] + lightweightCombination[lCom]["aero"] + gripCombination[gCom]["aero"] + toolsCombination[tCom]["aero"])*c["aero"],
                                "lightweight": (81+powerCombination[pCom]["lightweight"] + aeroCombination[aCom]["lightweight"] + lightweightCombination[lCom]["lightweight"] + gripCombination[gCom]["lightweight"] + toolsCombination[tCom]["lightweight"])*c["lightweight"],
                                "grip": (215+powerCombination[pCom]["grip"] + aeroCombination[aCom]["grip"] + lightweightCombination[lCom]["grip"] + gripCombination[gCom]["grip"] + toolsCombination[tCom]["grip"])*c["grip"],
                                "configSet": [powerCombination[pCom],aeroCombination[aCom],lightweightCombination[lCom],gripCombination[gCom],toolsCombination[tCom],c] 
                            }
                                config["totalPoints"] = config["power"] + config["aero"] + config["lightweight"] + config["grip"]
                                superCombination[pCom+"-"+aCom+"-"+lCom+"-"+gCom+"-"+tCom+"-"+str(c["id"])] = config    
                                totalPoints = config["power"] + config["aero"] + config["lightweight"] + config["grip"]
                                if config["cost"] <= 40 and totalPoints > maxPoints:
                                    maxPoints = totalPoints
                                    maxConfigSet = [powerCombination[pCom],aeroCombination[aCom],lightweightCombination[lCom],gripCombination[gCom],toolsCombination[tCom],c]
                                    maxConfig = config
                                
                                if config["cost"] <= 40 and config["power"] > 310 and config["aero"] > 310 and config["lightweight"] > 310 and config["grip"] > 310 and totalPoints > maxIdealPoints:
                                    maxIdealPoints = totalPoints
                                    maxIdealConfigSet = [powerCombination[pCom],aeroCombination[aCom],lightweightCombination[lCom],gripCombination[gCom],toolsCombination[tCom],c]
                                    maxIdealConfig = config

                                # combinations.append(config)

                                # if combiCount % 500000 == 0:
                                #     print("Insering %d combinations" % combiCount)
                                #     with open("data/"+str(combiCount)+".json", "w") as f:
                                #         json.dump(combinations, f)
                                #     # collection.insert_many(combinations)
                                #     combinations = []
                                #     print("Insered %d combinations" % combiCount)

    # if len(combinations) > 0:
    #     # collection.insert_many(combinations)
    #     with open("data/"+str(combiCount)+".json", "w") as f:
    #         json.dump(combinations, f)
    #     combinations = []
    #     print("Insered %d combinations" % combiCount)
    print("Max Points: ", maxPoints+(26+17+26))
    print("Max Config: ", maxConfig)
    print("Max Config Set: ", maxConfigSet)

    print("Max Ideal Points: ", maxIdealPoints+(26+17+26))
    print("Max Ideal Config: ", maxIdealConfig)
    print("Max Ideal Config Set: ", maxIdealConfigSet)

    print("Total Combinations:", combiCount)
    print("Time Taken: ", time.time()-timeNow)

    

if __name__ == "__main__":
    main()

# 31 273 273 272 293
# 2-3
# 2-3
# 3-5
# 7-8
# 1-2
# 1
# 1180

# 33 291 260 273 296
# 3-7
# 2-6
# 3-5
# 7-8
# 1-2
# 1
# 1189

# 36 325 316 286 317
# 3-4
# 2-3
# 3-5
# 4-8
# 1-2
# 2
# 1313
