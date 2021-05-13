@cnt = 0

def change(target, coins, usable)
  coin = coins.shift
  if coins.size == 0 then
    if target / coin <= usable then
      @cnt += 1
    end
  else
    # printf("n %d coin %d\n", target, coin)
    # p target/coin
    (0..target/coin).each{|i|
      change(target - coin * i, coins.clone, usable - i)
    }
  end
end

change(1000, [500, 100, 50, 10], 15)
puts @cnt


(0..0).each{|i|
  puts i
}