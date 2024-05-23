Problem statement: Design a limit order book

A limit order book contains prices and corresponding volumes with which people want to vuy a given stock

- Bid side represents open offers to buy
- Ask side represent open offers to sell
- Trades are made when highest bid &gte loewst ask (spread is crossed)
    - Spread is the difference between the lowest ask and the highest bid
    - When they cross, a trade is initiated because someone is willing to buy at prices someone is willing to sell at
    - Trade is initiated and volume is removed from the order book.
- If client submits a buy or sell order that cannot be filled, it gets stored in order book
- Orders are executed at the best possible price first, and if many, the earliest submitted is chosen

- Order Book Interface
    - PlaceOrder(order): Takes an order and either fills it or places it in the limit book, prints trades that  have taken place as a result

    - CancelOrder(orderId): takes an orderId and cancels it if has not yet been filled, otherwise is a nop

    - GetVolumeAtPrice(price, buyOrSellSide): Get volume of open orders for either buying or selling side of order book


- Order Class
    - orderId, timestamp, buyOrSellSide, price, volume, client
    - Corresponding getters and setters


- Data structures for Limit Order Book
- Process of submiting a buy order:
    - Check lowest price of the sell side of the limit book
    - If lowest price of seel side is &lte buy side, execute trade
    - If the buyer still has more volume left to fill, look at the next lowest price on the sell side and keep going
    - If there is unfilled volume for the buyres trade, add it to the buyer heap
- It seems a good case to use priority queues and heaps (min and max heaps). 
    - Sell side as a min heap will let us see the lowest sell price quickly
    - Buy side as a max heap will let us see the highest buy price quickly.
    - inserts and pops in (logn) every node of the heap
- Inside the heap
    - We can actually have multiple trades with the same price, so many orders could be at the same level on the heap
    - We'll need to use a queue for each node on our heap that tracks all the orders submitted in FIFO order.
        - Pulling from the beginning is the first order. 

- Getting the volume
    - We need to be able to get volume as fast as possible. 
    - Instead of ranging and acumulating the price on each level, we could maintain a hashmap that tracks volume of a given price. 
    - Each time we add or remove an order we update the volume map

- Cancellations
    - Actively cancel orders
        - Take the node representing the order and remove it from the heap
        - Removing arbitrary nodes from a q is O(n). We should use a queue via doubly linked list + hashmap so we can quickly locate node and remove in O(1)
        - If we removed all orders from that q, we need to remove the level from the q. We will need to map the q in the heap to to be able to do this quickly.
    - Lazy cancel orders
        - Mark orderes as cancelled and when we come upon it while trying to execute a trade we skip over.

Order Book Class
    - bestAsk minHeap: lowest sell
    - bestBid maxHeap: highest buy
    - orderMap[orderId to Order]: To cancel orders quickly.
    - volumeMap[price,buyOrSellSide to volume]: 
    - queueMap[price,buyOrSellSide to PriceQueue level]


- Place Order
    - Figure out opposite side of the book
    - Figure same side of the book
    - Map order ID to order.
    - While order has volume and opposite book isn't empty and while the first order to be filled in the opposite book( best price for us) has volume 
        - determine other order at top of other heap
        - determine trade price that will be the price of the other order alreay in book
        - determine volume that will be minimum between our order and the one in the book
        - Adjust our order volume and other order volume by subtracting the trade volume 
        - Print trade
        - If the other volume has no volume left, remove that order.
    - If we finished our while our order has volume left already, we will go and add our orderr to book
        - add Order To Book: 
            - If there's no price queue for given level, we need to create that queue and put the volume there
            - Else, we need to push the order to the existing queue and update volume

- Cancel Order
    - Look for order id in our order map
    - Get our price queue by using order price and order side
    - Remove our order from that price queue
    - If price queue len is 0
        - determine which book we are working with
        - remove price queue from the book (remove node from heap)
    - Update volume map subtracting order volume
    - delete order from order map


- GetVolumeAtPrice
    - Just return volume map given price and side. 0 if not exists.

- Extensions
    - Market orders
        - Just say buy or sell, don't specify price, just volume.
        - Go through opposite side of order book until volume for trade is exhausted and then do not store trade in limit order book
        