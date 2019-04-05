class BuildController < ApplicationController

  require 'rubygems'
  require 'zip'

  def index
  end

  def build_yml
    d="bar:\n - " + params[:url]
    d +="\n - " + params[:name]
    d +="\n - " + params[:from]
    d +="\n - " + params[:to]

    pathtoyml =  "#{Rails.root}/instance_info/instanceinfo.yml"
    folder =  "#{Rails.root}/instance_info/"
    input_filenames = ['instanceinfo.yml', 'instancecaller.go', 'Readme.txt']

    zipfile_name = "#{Rails.root}/archive.zip"
    #File.write(Rails.root + '/instanceinfo.yml', d)
    #file = "/home/orkun/instanceinfo.yml"
    File.open(pathtoyml, "w+") do |f|
      f.write(d)
    end

    Zip::File.open(zipfile_name, Zip::File::CREATE) do |zipfile|
      input_filenames.each do |filename|
        # Two arguments:
        # - The name of the file as it will appear in the archive
        # - The original file, including the path to find it
        zipfile.add(filename, File.join(folder, filename))
      end
      zipfile.get_output_stream("myFile") { |f| f.write "myFile contains just this" }
    end


#    send_file file, :filename => 'some.yml'
  end
end
