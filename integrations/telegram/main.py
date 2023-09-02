import telebot

f = open("./token", "r")
token = f.read()
f.close()

bot = telebot.TeleBot(token=token)

@bot.message_handler(func=lambda _: True)
def default(message):
    bot.send_message(message.chat.id, "Sorry, I don't understand that yet")