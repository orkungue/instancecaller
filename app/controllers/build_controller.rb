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
        zipfile.add(filename, File.join(folder, filename))
      end
    end
    zip_data = File.read(zipfile_name)
    send_data(zip_data, :type => 'application/zip', :filename => "filename.zip")
#    send_file file, :filename => 'some.yml'
  end
end
