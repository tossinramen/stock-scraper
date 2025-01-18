# Stock Scraper

## Overview
This Go program scrapes real-time stock data from Yahoo Finance for a predefined list of ticker symbols. It collects the following details for each stock:
- **Company**: The ticker symbol of the stock (e.g., `NVDA`, `TSLA`).
- **Price**: The current market price of the stock.
- **Percentage Change**: The percentage change in the stock's price for the day.
- **Price Change**: The absolute change in the stock's price for the day.

The scraped data is saved in a CSV file named `stocks.csv`.

---

## Features
- Scrapes data for multiple stock ticker symbols.
- Captures key stock metrics: company (ticker), price, percentage change, and price change.
- Saves the scraped data to a CSV file for easy analysis.
- Handles errors gracefully and logs them in the console.

---