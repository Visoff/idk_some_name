import telebot
import time

f = open("./token", "r")
token = f.read()
f.close()

bot = telebot.TeleBot(token=token)

@bot.message_handler(func=lambda _: True)
def default(message):
    bot.send_message(message.chat.id, "Sorry, I don't understand that yet")

@bot.inline_handler(func=lambda query: len(query.query) > 0)
def q(query):
    markup = telebot.types.InlineKeyboardMarkup()
    markup.add(telebot.types.InlineKeyboardButton("Button #1", callback_data="smth"))
    markup.add(telebot.types.InlineKeyboardButton("Button #2", callback_data="smth"))
    article = telebot.types.InlineQueryResultArticle(
            id=1,
            title=query.query,
            input_message_content=telebot.types.InputTextMessageContent(message_text=query.query),
            url="https://devsync.tech",
            reply_markup=markup
        )
    bot.answer_inline_query(query.id, [article])

bot.infinity_polling()