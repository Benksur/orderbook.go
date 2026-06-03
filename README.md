# orderbook.go

very naive orderbook implementation.

the orderbook class implements simple Add, Remove and Match methods, with orders sorted into two identical data structures based on their side (buy/sell). each side tracks orders by placing them into buckets mapped by their prices (levels), maintaining a sorted list of prices as well as mapping each orderID to their corresponding price. 
