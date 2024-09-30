a = Array.new(5, 'hello')
a[0].upcase!
puts a[0]
p a

b = Array.new(5) { 'hello' }
b[0].upcase!
puts b[0]
p b