CREATE TABLE `countrydata` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(150) NULL,
  `Region` VARCHAR(20) NULL,
  `Population` INT NULL,
  `Flag_png` VARCHAR(150) NULL,
  `Flag_svg` VARCHAR(150) NULL,
  `Currencies_Code` VARCHAR(25) NULL,
  `Currencies_Name` VARCHAR(50) NULL,
  `Currencies_Symbol` VARCHAR(25) NULL,
  PRIMARY KEY (`ID`));