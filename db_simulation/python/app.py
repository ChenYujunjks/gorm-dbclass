from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///test.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

class Customer(db.Model):
    __tablename__ = 'customers'
    email = db.Column(db.String(120), primary_key=True)
    name = db.Column(db.String(100))
    building_number = db.Column(db.String(50))
    street = db.Column(db.String(100))
    city = db.Column(db.String(100))
    state = db.Column(db.String(50))
    zip_code = db.Column(db.String(20))
    phones = db.relationship('CustomerPhone', backref='customer', cascade="all, delete-orphan", lazy=True)

class CustomerPhone(db.Model):
    __tablename__ = 'customer_phones'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    customer_email = db.Column(db.String(120), db.ForeignKey('customers.email'), nullable=False)
    phone = db.Column(db.String(20))

# 初始化数据库和插入数据
def init_db():
    db.create_all()

    # 插入数据
    customer1 = Customer(
        email="alice@example.com",
        name="Alice Wang",
        building_number="123",
        street="Main St",
        city="NY",
        state="NY",
        zip_code="10001",
        phones=[CustomerPhone(phone="123-456-7890"), CustomerPhone(phone="098-765-4321")]
    )

    customer2 = Customer(
        email="bob@example.com",
        name="Bob Li",
        building_number="456",
        street="Market St",
        city="SF",
        state="CA",
        zip_code="94105",
        phones=[CustomerPhone(phone="555-555-5555")]
    )

    db.session.add(customer1)
    db.session.add(customer2)
    db.session.commit()

@app.route('/customers')
def get_customers():
    # 联接查询
    results = db.session.query(Customer.name, CustomerPhone.phone).join(CustomerPhone).all()

    # 输出结果
    output = []
    for result in results:
        output.append(f"{result.name}: {result.phone}")

    return "<br>".join(output)

if __name__ == '__main__':
    # 初始化数据库并插入数据
    with app.app_context():
        init_db()
    
    app.run(debug=True)
