
{
	description = "";
	
	inputs = {
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38.url = "github:nixos/nixpkgs/6d974716629d3bf8a31b519952c6a5754a41bd38";
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14.url = "github:nixos/nixpkgs/ac5c1886fd9fe49748d7ab80accc4c847481df14";
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7.url = "github:nixos/nixpkgs/1ebb7d7bba2953a4223956cfb5f068b0095f84a7";
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746.url = "github:nixos/nixpkgs/a731d0cb71c58f56895f71a5b02eda2962a46746";
			
		nixpkgs.url = "github:nixos/nixpkgs/nixos-23.11";
	};
	
	outputs = { self, nixpkgs,  nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38, 
	 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14, 
	 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7, 
	 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746, 
	 }: let
	  supportedSystems = [ "x86_64-linux" "aarch64-darwin" "x86_64-darwin" "aarch64-linux" ];
	  forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs = import nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38 { inherit system; };
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs = import nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14 { inherit system; };
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs = import nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7 { inherit system; };
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs = import nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746 { inherit system; };
		
		pkgs = import nixpkgs { inherit system; };
	  });
	in {
	  packages = forEachSupportedSystem ({ pkgs,  nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs, 
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs, 
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs, 
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs, 
		 }: {
		default = pkgs.callPackage ./default.nix {
		  
		};
	  });
	
	  devShells = forEachSupportedSystem ({ pkgs,  nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs, 
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs, 
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs, 
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs, 
		 }: {
		devShell = pkgs.mkShell {
		  # The Nix packages provided in the environment
		  packages =  [
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.api-linter  
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.buf  
			nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs.delve  
			nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs.go  
			nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs.gotools  
			
		  ];
		};
	  });
	
	  runtimeEnvs = forEachSupportedSystem ({ pkgs,  nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs,  nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs,  nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs,  nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs,  }: {
		runtime = pkgs.buildEnv {
		  name = "runtimeenv";
		  paths = [ 
			nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs.cacert   
			
		   ];
		};
	});
	};
}
