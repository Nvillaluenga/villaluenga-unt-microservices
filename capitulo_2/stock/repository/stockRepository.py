from random import randint, random
from tinydb import Query, TinyDB
from tinydb.table import Document 

stock_table = TinyDB('stock.json').table('stock')



def get_stock(product_id: int) -> dict:
  stock = stock_table.get(doc_id=product_id)
  return stock.copy()

def update_stock(product_id: int, stock: int):
  stock_table = TinyDB('stock.json').table('stock')
  stock_table.update({"stock":stock}, doc_ids = [product_id])

if __name__ == "__main__":
  for i in range(100):
    stock = randint(1, 15)
    price = round((random() * 100), 2)
    stock_table.insert(Document({"stock":stock, "price": price}, doc_id=i))
  # print(get_stock(15))