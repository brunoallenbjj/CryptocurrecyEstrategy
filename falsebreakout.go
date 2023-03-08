


//@version=5
strategy("False Breakout", overlay=true)

// Define input variables

var int tradeSize = input.int(100, "Trade size", minval=1)
var float stopLossPct = input.float(1,"Stop Loss %", minval=0.1, step=0.1)
var float takeProfitPct = input.float(2, "Take Profit%", minval=0.1, stop=0.1)
var bool buySignal = false
var bool sellSignal = false

// Define false breakout condition

var bool falseBreakout = false
var float lastHigh = ta.high[1]
var float lastLow = ta.low[1]
var float currentHigh = ta.high
var float currentLow = ta.low

if strategy.position_size == 0 
    if currentHigh > lastHigh and currentLow > lastLow
	    falseBreakout := true
	    buySignal := falseBreakout

	else if currentHigh < lastHigh and currentLow < lastLow
	    falseBreakout := true
		sellSignal := falseBreakout

else if strategy.position_size	> 0
    if currentLow < lastLow
	    falseBreakout := true 
		strategy.exit("Ëxit", "Long")

else if strategy.position_size	< 0 
    if currentHigh > lastHigh 
	    falseBreakout := true 
		strategy.exit("Ëxit", "Short")	 	


	
// Execute Trades

if buySignal 

    strategy.entry("Long", strategy.long, qty=tradeSize, comment="Long")
	strategy.exit("Stop Loss", "Long", loss=stopLossPct/100)
	strategy.exit("Take Profit", "Long", profit.takeProfitPct/100)

if	sellSignal
    strategy.entry("Short", strategy.short, qty=tradeSize, comment="Short")
	strategy.exit("Stop Loss", "Short", loss=stopLossPct/100)
	strategy.exit("Take Profit", "Short", profit.takeProfitPct/100)

	// Plot buy and sell signals

	plotshape(buySignal,style=shape.triangleup, color=color.green, location=location.belowbar, size=size.small, title="Buy Signal")
	plotshape(sellSignal, style=shape.triangledown, color=color.red, location=location.abovebar, size=size.small, title="Sell Signal")
	