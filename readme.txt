###################################
###   automatic fill in robot   ###
###################################

## INTRO

This program is specially designed for the program TaxationOfflineClient on a 1920x1080 screen resolution.
For all othe programs or screen resolutions the program has to be adopted.


## DATA

Input data in the form of csv-files have to be placed into the "data" directory.
The files can be produced by scripts avalible in github.com/satlawa/obf_autman.
The naming convention is:

    autax_x.csv   where x can be: "wo", "text", "nutz", "bz"

    exp.: autax_bz.csv

Example files are provided in the directory "data".	


## RUN

For the program to work correctly the TaxationOfflineClient has to be running in the foreground in a maximised window size.
Additonally the the x wo-instance has to be maked, where x is the starting page.
Than the program has to be executed from the command line:

    $ autfillin_x y z    where x can be: "wo", "text", "nutz", "bz"; y is the start page number; z is the end page number

    exp.: autfillin_text 0 100   this will run the program filling in the "text" data from page 0 to 100

it is advisible to use the following sequence to fill in the data: 

   "wo" -> "text" -> "nutz" -> "bz"

p.s.: the program autfillin_bz takes 3 arguments where the third argument is taking "y" for yes or "n" for no to create the
longes "Auswertekategorie" in order to prevent the change the location of the buttons.

   exp.: autfillin_bz 0 100 y
 
