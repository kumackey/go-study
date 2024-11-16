require 'aws-sdk-s3'

s3 = Aws::S3::Client.new

bucket = 'kumackey-sample-test-20241116'
body = 'Hello, World!'
output = s3.put_object(bucket: bucket, key: 'sample.txt', body: body)
etag = output.etag.gsub(/"/, '')
puts "Etag: #{etag}"

begin
  s3.put_object(bucket: bucket, key: 'sample.txt', body: body, if_none_match: etag)
rescue Aws::S3::Errors::ServiceError => e
  puts e.class
  puts e.message
  puts e.context.http_request.headers.inspect
end

puts '-----------------'

begin
  s3.put_object(bucket: bucket, key: 'sample.txt', body: body, if_none_match: '*')
rescue Aws::S3::Errors::ServiceError => e
  puts e.class
  puts e.message
  puts e.context.http_request.headers.inspect
end
