from flask import Flask, request, jsonify
from http import HTTPStatus
from repository.stockRepository import get_stock, update_stock

app = Flask(__name__)

@app.route('/stock/diminish/<int:product_id>/<int:ammount>', methods=['POST'])
def diminish_stock_handler(product_id: int, ammount: int):
    # Call the stock function
    stock = get_stock(product_id).get("stock")
    new_ammount = stock - ammount
    if (new_ammount < 0):
      return jsonify(product_id=product_id, stock=-1), HTTPStatus.CONFLICT
    else:
      update_stock(product_id=product_id, stock=new_ammount)
      return jsonify(product_id=product_id, stock=new_ammount), HTTPStatus.OK


@app.route('/stock/<int:product_id>', methods=['GET'])
def get_stock_handler(product_id: int):
    # Call the stock function
    stock = get_stock(product_id)

    return jsonify(product_id=product_id, stock=stock["stock"], price=stock["price"]), HTTPStatus.OK

if __name__ == '__main__':
    app.run(debug=True, port=8081)
