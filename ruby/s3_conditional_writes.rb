require 'aws-sdk-s3'

s3 = Aws::S3::Client.new
# aws s3 ls

bucket = ENV.fetch('S3_BUCKET')

puts '【準備1】まずはファイルをアップロード'
body = 'Hello, World!'
s3.put_object(bucket: bucket, key: 'sample.txt', body: body)

puts '-----------------'
puts '【準備2】アップロードできたことと、Etagと中身の確認'
output = s3.get_object(bucket: bucket, key: 'sample.txt')

# Etagはダブルクォートで囲まれているので、取り除く
etag = output.etag.gsub(/"/, '')
puts "Etag: #{etag}"
puts "Body: #{output.body.read}"

puts '-----------------'
puts '【動作確認1】put_objectのconditional writes(ETagが一致する場合)'
begin
  s3.put_object(bucket: bucket, key: 'sample.txt', body: body, if_none_match: etag)
rescue Aws::S3::Errors::PreconditionFailed => e
  puts e.message
  puts 'Aws::S3::Errors::PreconditionFailedがでてアップロードが失敗'
end

puts '-----------------'
puts '【動作確認2】put_objectのconditional writes(ETagが一致しない場合)'
output = s3.put_object(bucket: bucket, key: 'sample.txt', body: body, if_none_match:'different_etag')
puts "Etag: #{output.etag}"
puts 'Aws::S3::Errors::PreconditionFailedがでないのでアップロードが成功'

puts '-----------------'
puts '【準備3】コピー先の作成'
output = s3.copy_object(bucket: bucket, key: 'sample_target.txt', copy_source: "#{bucket}/sample.txt")
puts "Last modified: #{output.copy_object_result.last_modified}"

puts '-----------------'
puts '【動作確認3】copy_objectのconditional writes(ETagが一致する場合)'
begin
  # 引数名がif_none_matchでなくcopy_source_if_none_matchであることに注意！これで1時間くらいハマった・・・
  s3.copy_object(bucket: bucket, key: 'sample_target.txt', copy_source: "#{bucket}/sample.txt", copy_source_if_none_match: etag)
rescue Aws::S3::Errors::PreconditionFailed => e
  puts e.message
  puts 'Aws::S3::Errors::PreconditionFailedがでてコピーが失敗'
end

puts '-----------------'
puts '【動作確認4】copy_objectのconditional writes(ETagが一致しない場合)'
s3.copy_object(bucket: bucket, key: 'sample_target.txt', copy_source: "#{bucket}/sample.txt", copy_source_if_none_match:'different_etag')
puts 'Aws::S3::Errors::PreconditionFailedがでないのでコピーが成功'