from random import randint
from tinydb import TinyDB
from tinydb.table import Document 

promotion_product_table = TinyDB('promotion.json').table('promotion_product')


def get_prod_promotion(product_id: int) -> dict:
  promotion = promotion_product_table.get(doc_id=product_id)
  return promotion.copy()

def get_prod_promotions() -> list[dict]:
  promotions = promotion_product_table.all()
  promotions = [p.copy() for p in promotions]
  return promotions


if __name__ == "__main__":
  for i in range(100):
    promotion = randint(-15, 15)
    if promotion > 5:
      promotion_product_table.insert(Document({"promotion":promotion, "product_id": i}, doc_id=i))
  # print(get_stock(15))