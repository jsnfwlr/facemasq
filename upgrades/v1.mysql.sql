CREATE TABLE Categories (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    Icon VARCHAR(64)NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE Status (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    Icon VARCHAR(64)NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE DeviceTypes (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    Icon VARCHAR(64)NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE InterfaceTypes (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    Icon VARCHAR(64)NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE VLANs (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    IPv4Mask TEXT NOT NULL,
    IPv6Mask TEXT NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE Locations (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    IsCloud BOOL DEFAULT FALSE NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

CREATE TABLE OperatingSystems (
    ID INT NOT NULL AUTO_INCREMENT,
    Vendor VARCHAR(128) NOT NULL,
    Family VARCHAR(128) NOT NULL,
    Version VARCHAR(128) NOT NULL,
    Name VARCHAR(128) NOT NULL,
    IsOpenSource BOOL DEFAULT FALSE NOT NULL,
    IsServer BOOL DEFAULT FALSE NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    CONSTRAINT UC_OperatingSystem UNIQUE (Vendor, Family, Version, Name)
);

CREATE TABLE Architectures (
    ID INT NOT NULL AUTO_INCREMENT,
    Label VARCHAR(128) NOT NULL,
    BitSpace INT DEFAULT 64 NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Label)
);

-- UserData   - Users and maintainer records: users can log in to the app, maintainers are responsible for devices
CREATE TABLE Users (
    ID INT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(128),
    Password VARCHAR(128),
    Label VARCHAR(128) NOT NULL,
    CanAuthenticate BOOL DEFAULT FALSE NOT NULL,
    AccessLevel INT DEFAULT 0 NOT NULL,
    IsInternal BOOL DEFAULT FALSE NOT NULL,
    Notes LONGTEXT,
    IsLocked BOOL DEFAULT FALSE NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Username),
    UNIQUE (Label)
);

-- MetaData   - App settings & User preferences
CREATE TABLE Meta (
    Name VARCHAR(128) NOT NULL,
    Value VARCHAR(128) NOT NULL,
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users( ID),
    CONSTRAINT UC_Meta UNIQUE (Name, UserID)
);

-- DeviceData - Tables that define the devices, their interfaces, addresses, and hostnames
CREATE TABLE Devices (
    ID INT NOT NULL AUTO_INCREMENT,
    MachineName VARCHAR(128) DEFAULT "Unknown" NOT NULL,
    Brand VARCHAR(128),
    Model VARCHAR(128),
    Purchased DATETIME,
    Serial VARCHAR(128),
    IsTracked BOOL DEFAULT FALSE NOT NULL,
    FirstSeen DATETIME NOT NULL,
    IsGuest BOOL DEFAULT FALSE NOT NULL,
    IsOnline BOOL DEFAULT FALSE NOT NULL,
    Label VARCHAR(128),
    Notes LONGTEXT,
    CategoryID INT DEFAULT 1 NOT NULL,
    StatusID INT DEFAULT 1 NOT NULL,
    MaintainerID INT DEFAULT 1 NOT NULL,
    LocationID INT DEFAULT 1 NOT NULL,
    DeviceTypeID INT DEFAULT 1 NOT NULL,
    OperatingSystemID INT DEFAULT 1 NOT NULL,
    ArchitectureID INT DEFAULT 1 NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(ID),
    FOREIGN KEY (StatusID) REFERENCES Status(ID),
    FOREIGN KEY (MaintainerID) REFERENCES Users(ID),
    FOREIGN KEY (LocationID) REFERENCES Locations(ID),
    FOREIGN KEY (DeviceTypeID) REFERENCES DeviceTypes(ID),
    FOREIGN KEY (OperatingSystemID) REFERENCES OperatingSystems(ID),
    FOREIGN KEY (ArchitectureID) REFERENCES Architectures(ID)
);

CREATE TABLE Interfaces (
    ID INT NOT NULL AUTO_INCREMENT,
    MAC VARCHAR(17) NOT NULL,
    IsPrimary BOOL DEFAULT TRUE NOT NULL,
    IsVirtual BOOL DEFAULT FALSE NOT NULL,
    IsOnline BOOL DEFAULT FALSE NOT NULL,
    Label VARCHAR(128),
    Notes LONGTEXT,
    LastSeen DATETIME NOT NULL,
    StatusID INT DEFAULT 1 NOT NULL,
    InterfaceTypeID INT DEFAULT 1 NOT NULL,
    VLANID INT DEFAULT 1 NOT NULL,
    DeviceID INT NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (MAC),
    FOREIGN KEY (StatusID) REFERENCES Status(ID),
    FOREIGN KEY (InterfaceTypeID) REFERENCES InterfaceTypes(ID),
    FOREIGN KEY (VLANID) REFERENCES VLANs(ID),
    FOREIGN KEY (DeviceID) REFERENCES Devices(ID)
);

CREATE TABLE Addresses (
    ID INT NOT NULL AUTO_INCREMENT,
    IPv4 VARCHAR(15),
    IPv6 VARCHAR(128),
    IsPrimary BOOL DEFAULT TRUE NOT NULL,
    IsVirtual BOOL DEFAULT FALSE NOT NULL,
    IsReserved BOOL DEFAULT FALSE NOT NULL,
    LastSeen DATETIME NOT NULL,
    Label VARCHAR(128),
    Notes LONGTEXT,
    InterfaceID INT NOT NULL,
    PRIMARY KEY (ID),
    CONSTRAINT UC_Address UNIQUE (InterfaceID,IPv4),
    FOREIGN KEY (InterfaceID) REFERENCES Interfaces(ID)
);

CREATE TABLE Hostnames (
    ID INT NOT NULL AUTO_INCREMENT,
    Hostname VARCHAR(512) NOT NULL,
    IsDNS BOOL DEFAULT FALSE NOT NULL,
    IsSelfSet BOOL DEFAULT FALSE NOT NULL,
    Notes LONGTEXT,
    AddressID INT NOT NULL,
    PRIMARY KEY (ID),
    UNIQUE (Hostname),
    FOREIGN KEY (AddressID) REFERENCES Addresses(ID)
);

-- ScanData   - Tables that record the times that addresses were online
CREATE TABLE Scans (
    ID INT NOT NULL AUTO_INCREMENT,
    Time DATETIME NOT NULL UNIQUE,
    PRIMARY KEY (ID)
);

CREATE TABLE History (
    AddressID INT NOT NULL,
    ScanID INT NOT NULL,
    CONSTRAINT UC_History UNIQUE(AddressID,ScanID),
    FOREIGN KEY (AddressID) REFERENCES Addresses(ID),
    FOREIGN KEY (ScanID) REFERENCES Scans(ID)
);

CREATE TABLE Ports (
    AddressID INT NOT NULL,
    ScanID INT NOT NULL,
    Protocol VARCHAR(8) NOT NULL,
    Port INT NOT NULL,
    CONSTRAINT UC_Port UNIQUE (AddressID,ScanID,Protocol,Port),
    FOREIGN KEY (AddressID) REFERENCES Addresses(ID),
    FOREIGN KEY (ScanID) REFERENCES Scans(ID)
);


INSERT INTO Categories (Label, Icon, Notes, IsLocked) 
  VALUES
    ("Unsorted", "HelpCircle", NULL, 1);

INSERT INTO Status (Label, Icon, Notes, IsLocked)
  VALUES 
    ("Invading", "HelpCircle", NULL, 1);

INSERT INTO DeviceTypes (Label, Icon, Notes, IsLocked) 
  VALUES
    ("Unspecified", "HelpCircle", NULL, 1);

INSERT INTO InterfaceTypes (Label, Icon, Notes, IsLocked)
  VALUES 
    ("WiFi", "HelpCircle", NULL, 1),
    ("Ethernet Cable", "HelpCircle", NULL, 1),
    ("Fibre", "HelpCircle", NULL, 1),
    ("Internal", "HelpCircle", "For MACVLAN containers, Virtual interfaces, etc", 1);

INSERT INTO VLANs (Label, Notes, IPv4Mask, IPv6Mask, IsLocked) 
  VALUES 
    ("Default", NULL, "0.0.0.0", "", 1);

INSERT INTO Locations (Label, IsCloud, Notes, IsLocked)
  VALUES
    ("Limbo", 0, NULL, 1);

INSERT INTO OperatingSystems (Vendor, Family, Version, Name, IsOpenSource, IsServer, Notes, IsLocked)
  VALUES 
    ("?", "?", "?", "?", 0, 0, NULL, 1),
    ("Apple", "MacOS", "10.13", "High Sierra", 0, 0, NULL, 0),
    ("Apple", "MacOS", "10.14", "Mojave", 0, 0, NULL, 0),
    ("Apple", "MacOS", "10.15", "Catalina", 0, 0, NULL, 0),
    ("Apple", "MacOS", "11", "Big Sur", 0, 0, NULL, 0),
    ("Apple", "MacOS", "12", "Monterey", 0, 0, NULL, 0),
    ("Microsoft", "Windows", "10", "Win10", 0, 0, NULL, 0),
    ("Microsoft", "Windows", "11", "Win11", 0, 0, NULL, 0),
    ("Canonical", "Ubuntu", "18.04", "Bionic Beaver", 1, 0, NULL, 0),
    ("Canonical", "Ubuntu", "20.04", "Focal Fossa", 1, 0, NULL, 0),
    ("Canonical", "Ubuntu", "21.04", "Hirsute Hippo", 1, 0, NULL, 0),
    ("Canonical", "Ubuntu", "21.10", "Impish Indri", 1, 0, NULL, 0),
    ("Canonical", "Ubuntu", "22.04", "Jammy Jellyfish", 1, 0, NULL, 0),
    ("Google", "Android", "6", "Marshmallow", 1, 0, NULL, 0),
    ("Google", "Android", "7", "Nougat", 1, 0, NULL, 0),
    ("Google", "Android", "8", "Oreo", 1, 0, NULL, 0),
    ("Google", "Android", "9", "Pie", 1, 0, NULL, 0),
    ("Google", "Android", "10", "Q", 1, 0, NULL, 0),
    ("Google", "Android", "11", "R", 1, 0, NULL, 0),
    ("Google", "Android", "12", "S", 1, 0, NULL, 0),
    ("Google", "Android", "13", "T", 1, 0, NULL, 0);

INSERT INTO Architectures (Label, BitSpace, Notes, IsLocked)
  VALUES
    ("Unknown", 0, NULL, 1),
    ("x86", 32, NULL, 0),
    ("x64", 64, NULL, 0),
    ("ARM", 32, NULL, 0),
    ("ARM64", 64, NULL, 0),
    ("RISC-V 32I", 32, NULL, 0),
    ("RISC-V 32E", 32, NULL, 0),
    ("RISC-V 64I", 64, NULL, 0),
    ("RISC-V 128I", 128, NULL, 0);

INSERT INTO Users (Username, Password, Label, CanAuthenticate, AccessLevel, IsInternal, Notes, IsLocked)
  VALUES
    (NULL, NULL, "Invader", 0, 0, 0, NULL, 1),
    ("Admin", "{{ .PasswordHash }}", "Admin", 1, 1, 0, NULL, 1);

INSERT INTO Meta (Name, Value)
  VALUES 
    ("DBVersion", "1");
