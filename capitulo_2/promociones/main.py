from flask import Flask, jsonify
from http import HTTPStatus
from repository.promotionsRepository import get_prod_promotion, get_prod_promotions

app = Flask(__name__)

@app.route('/promotions/product/<int:product_id>', methods=['GET'])
def diminish_stock_handler(product_id: int):
    # Call the stock function
    promotion = get_prod_promotion(product_id)
    return jsonify(promotion), HTTPStatus.OK


@app.route('/promotions/product', methods=['GET'])
def get_stock_handler():
    # Call the stock function
    promotions = get_prod_promotions()
    return jsonify(promotions=promotions), HTTPStatus.OK


if __name__ == '__main__':
    app.run(debug=True, port=8082)
