for i in 1..30
  if i % 15 == 0
    puts "FizzBuzz"
  else
    if i % 5 == 0
      puts "Buzz"
    else
      if i % 3 == 0
        puts "Fizz"
      else
        puts i
      end
    end
  end
end

# 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, ...