require 'faraday'
require 'byebug'

Faraday.post('http://localhost:8000', { name: 'Alice' }.to_json)
# (byebug) request
# #<struct Faraday::Request http_method=:post, path="http://localhost:8000", params={}, headers={"User-Agent"=>"Faraday v2.7.10"}, body="{\"name\":\"Alice\"}", options=#<Faraday::RequestOptions (empty)>>

Faraday.delete('http://localhost:8000', { name: 'Alice' }.to_json)
# ..../lib/faraday/utils/params_hash.rb:28:in `update': undefined method `each' for "{\\"name\\":\\"Alice\\"}":String (NoMethodError)

Faraday.delete('http://localhost:8000', { name: 'Alice' })
# (byebug) request
# #<struct Faraday::Request http_method=:delete, path="http://localhost:8000", params={"name"=>"Alice"}, headers={"User-Agent"=>"Faraday v2.7.10"}, body=nil, options=#<Faraday::RequestOptions (empty)>>
# (byebug) exclusive_url
# #<URI::HTTP http://localhost:8000?name=Alice>

Faraday.delete('http://localhost:8000', { name: 'Alice' }) do |req|
  req.body = { name: 'Alice' }.to_json
end
# (byebug) request
# #<struct Faraday::Request http_method=:delete, path="http://localhost:8000", params={"name"=>"Alice"}, headers={"User-Agent"=>"Faraday v2.7.10"}, body="{\"name\":\"Alice\"}", options=#<Faraday::RequestOptions (empty)>>