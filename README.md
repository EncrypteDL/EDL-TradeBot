# **EDL TradeBot**

**EDL TradeBot** is an advanced automated trading bot developed using Golang. Designed to operate across multiple cryptocurrency exchanges, it supports a variety of trading strategies and offers robust features for seamless asset trading. The bot integrates key functionalities like risk management, custom strategies, and API integration with top exchanges.

---

## **⚠️ Project Status**

**This project is currently in active development.** Some features may be incomplete or subject to change, and stability is not guaranteed. Use it at your own risk during this development phase.

---

## **Features**
- **Multi-Exchange Support**: Integrated with Binance, Coinbase, Kraken, Bybit, Pionex, and more.
- **Algorithmic Trading**: Implement various customizable strategies to optimize trades.
- **Risk Management**: Includes stop-loss, take-profit, and position-sizing features.
- **Backtesting**: Test strategies against historical data before live deployment.
- **Modular Architecture**: Clean, layered structure for easy scalability and maintenance.
- **24/7 Trading**: Automates trading decisions and operates without interruption.

---

## **Project Structure**

The project is organized into the following main packages and layers:

```bash
EDL-TradeBot/
├── cmd/                  # Main application entry points
├── exchange/             # Exchange integrations
│   ├── binance/
│   ├── coinbase/
│   ├── kraken/
│   ├── bybit/
│   ├── pionex/
│   └── ...               # Add more exchanges as needed
├── internal/             
│   ├── database/         # Database management and storage
│   ├── backend/          # Backend services and utilities
│   ├── config/           # Configuration management
│   └── ...               # Additional internal utilities
├── money/                # Money management, account balances, and handling
├── strategies/           # Custom trading strategies
│   ├── arbitrage/
│   ├── market_making/
│   ├── trend_following/
│   └── ...               # Add more strategies as needed
└── README.md             # Project documentation
```

---

## **Installation**

To run the EDL TradeBot, ensure that you have Golang installed on your system.

1. **Clone the repository**:
   ```bash
   git clone https://github.com/EncrypteDL/EDL-TradeBot.git
   ```
2. **Navigate to the project directory**:
   ```bash
   cd EDL-TradeBot
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

---

## **Configuration**

1. **Exchange API Keys**:
   Each exchange requires its own set of API keys for trading. Place your API keys in a configuration file located under `internal/config/`:
   ```yaml
   binance:
     apiKey: "your-binance-api-key"
     secretKey: "your-binance-secret-key"
   coinbase:
     apiKey: "your-coinbase-api-key"
     secretKey: "your-coinbase-secret-key"
   ```

2. **Database Configuration**:
   The bot requires a database for storing trade history, user data, and strategies. Configure your database connection in the `internal/database/` package:
   ```yaml
   database:
     type: "postgres"       # Supported types: postgres, mysql, sqlite
     user: "db_user"
     password: "db_password"
     dbname: "trading_db"
     host: "localhost"
     port: 5432
   ```

3. **Strategies Configuration**:
   You can configure the trading strategies in the `strategies/` package. Each strategy has customizable parameters based on your trading preferences.

---

## **Usage**

### **Running the Bot**
To run the trading bot, use the following command:
```bash
go run cmd/main.go
```

### **Backtesting**
You can backtest a strategy using historical market data before deploying it live:
```bash
go run cmd/backtest.go --strategy=trend_following --exchange=binance
```

---

## **Supported Exchanges**
EDL TradeBot currently supports the following exchanges:
- **Binance**
- **Coinbase**
- **Kraken**
- **Bybit**
- **Pionex**

More exchanges can be added as required.

---

## **Custom Strategies**
The bot allows you to develop custom trading strategies and configure them in the `strategies/` directory. Here are some of the provided strategies:

1. **Arbitrage**:
   - Exploits price differences across multiple exchanges.
   
2. **Market Making**:
   - Provides liquidity by placing buy and sell orders to profit from the spread.

3. **Trend Following**:
   - Follows market trends and places trades based on momentum indicators.

---

## **Contributing**

We welcome contributions! Feel free to open issues, submit pull requests, or suggest new features. Ensure that your contributions align with the modular architecture of the bot.

### **Development Setup**

1. Fork the repository.
2. Create a new branch.
3. Make changes and commit them.
4. Open a pull request for review.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

## **Contact**

For more information, contact the project lead [Zakaria Saif](mailto:zakaria@encrypteDL.com).
