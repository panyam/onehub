import ipdb
import random
import csv
import requests

def ensure_users():
    out = []
    auth = f"admin:admin123"
    for i in range(100):
        userid = f"ltuser{i}"
        resp = requests.get(f"http://{auth}@localhost:7080/api/v1/users/{userid}")
        if resp.status_code == 200:
            user = resp.json()["user"]
            out.append(user)
            print("Found User: ", user)
        else:
            # create it
            user = { "id": userid, "name": make_random_name() }
            resp = requests.post(f"http://{auth}@localhost:7080/api/v1/users", json = {"user": user})
            user = resp.json()["user"]
            out.append(user)
            print("Created User: ", user)
    return out

def sendmsg(uid, tid, msg):
    payload = {"message": {
        "topic_id": tid,
        "user_id": uid,
        "content_type": "text/plain",
        "content_text": msg,
    }}
    auth = f"{uid}:{uid}123"
    return requests.post(f"http://{auth}@localhost:7080/api/v1/{currtid}/messages", json= payload)["message"]

def ensure_topics():
    users = ensure_users()
    lines = list(csv.reader(open("./chatmessages.csv")))
    topics = {}
    currtid = None
    for tid, msg in lines:
        creator = random.choice(users)["id"]
        auth = f"{creator}:{creator}123"
        if tid != currtid:
            currtid = tid
            # create new topic
            tid = f"lt{tid}"
            topicname = msg
            topic = {"topic": { "id": tid, "name": topicname, }}
            resp = requests.get(f"http://{auth}@localhost:7080/api/v1/topics/{tid}")
            if resp.status_code == 200:
                topic = resp.json()["topic"]
                print("Found Topic: ", topic)
            else:
                resp = requests.post(f"http://{auth}@localhost:7080/api/v1/topics", json = topic)
                topic = resp.json()["topic"]
                print("Created Topic: ", topic)
            topics[topic["id"]] = topic
    return users, topics

def generate_messages():
    users, topics = ensure_topics()
    lines = list(csv.reader(open("./chatmessages.csv")))
    for tid, msg in lines:
        creator = random.choice(users)["id"]
        auth = f"{creator}:{creator}123"
        sendmsg(creator, f"lt{tid}", msg)

def make_random_name():
  adj = random.choice(ADJECTIVES)
  animal = random.choice(ANIMALS)
  adj = adj[0].upper() + adj[1:]
  animal = animal[0].upper() + animal[1:]
  return f"{adj} {animal}"

ADJECTIVES = [
  "adaptable",
  "adventurous",
  "affable",
  "affectionate",
  "agreeable",
  "ambitious",
  "amiable",
  "amicable",
  "amusing",
  "brave",
  "bright",
  "broadminded",
  "calm",
  "careful",
  "charming",
  "communicative",
  "compassionate	",
  "conscientious",
  "considerate",
  "convivial",
  "courageous",
  "courteous",
  "creative",
  "decisive",
  "determined",
  "diligent",
  "diplomatic",
  "discreet",
  "dynamic",
  "easygoing",
  "emotional",
  "energetic",
  "enthusiastic",
  "exuberant",
  "fair-minded",
  "faithful",
  "fearless",
  "forceful",
  "frank",
  "friendly",
  "funny",
  "generous",
  "gentle",
  "good",
  "gregarious",
  "hard-working",
  "helpful",
  "honest",
  "humorous",
  "imaginative",
  "impartial",
  "independent",
  "intellectual",
  "intelligent",
  "intuitive",
  "inventive",
  "kind",
  "loving",
  "loyal",
  "modest",
  "neat",
  "nice",
  "optimistic",
  "passionate",
  "patient",
  "persistent	",
  "pioneering",
  "philosophical",
  "placid",
  "plucky",
  "polite",
  "powerful",
  "practical",
  "pro-active",
  "quick-witted",
  "quiet",
  "rational",
  "reliable",
  "reserved",
  "resourceful",
  "romantic",
  "self-confident",
  "self-disciplined",
  "sensible",
  "sensitive",
  "shy",
  "sincere",
  "sociable",
  "straightforward",
  "sympathetic",
  "thoughtful",
  "tidy",
  "tough",
  "unassuming",
  "understanding",
  "versatile",
  "warmhearted",
  "willing",
  "witty",
]

ANIMALS = [
  'Tiger',
  'Lion',
  'Elephant',
  'Leopard',
  'Panther',
  'Cheetah',
  'Wolf',
  'Jaguar',
  'Hyena',
  'Giraffe',
  'Deer',
  'Zebra',
  'Gorilla',
  'Monkey',
  'Chimpanzee',
  'Bear',
  'Wild Boar',
  'Hippopotamus',
  'Kangaroo',
  'Rhinoceros',
  'Crocodile',
  'Panda',
  'Squirrel',
  'Mongoose',
  'Porcupine',
  'Koala Bear',
  'Wombat',
  'Meerkat',
  'Otter',
  'Hedgehog',
  'Possum',
  'Chipmunk',
  'Squirrel',
  'Raccoon',
  'Jackal',
  'Hare',
  'Mole',
  'Rabbit',
  'Alligator',
  'Monitor Lizard',
  'Oryx',
  'Elk',
  'Badger',
  'Dinosaur',
  'Pangolin',
  'Mole',
  'Okapi',
  'Camel',
  'Wild cat',
  'Coyote',
  'Bison',
  'African Elephant',
  'Aardvark',
  'Antelope',
  'Alpine Goat',
  'Komodo Dragon',
  'Bearded Dragon',
  'Royal Bengal Tiger',
  'Flying Squirrel',
  'Emu',
  'Eagle',
  'Eel',
  'Asiatic Lion',
  'Armadillo',
  'Beaver',
  'Emperor Penguin',
  'Baboon',
  'Bat',
  'Chameleon',
  'Bull',
  'Giant Panda',
  'Chihuahua',
  'Orangutan',
  'Chinchillas',
  'Hawk',
  'Iguana',
  'Ibis',
  'Ibex',
  'King Cobra',
  'Jellyfish',
  'Goose',
  'Walrus',
  'Seal',
  'Skink',
  'Markhor',
  'Falcon',
  'Bull Shark',
  'Arctic Wolf',
  'Owl',
  'Bulbul',
  'Bobcat',
  'Guinea Pig',
  'Yak',
  'Reindeer',
  'Moose',
  'Puma',
  'Okapi',
  'Marten',
  'Squirrel Monkey',
  'Caracal'
]

