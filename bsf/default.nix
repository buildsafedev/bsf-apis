
	{
	   lib,
	   stdenv,
	   buildGoModule,
	   ... 
	 }: buildGoModule {
	   name = "bsf-apis";
	   src = ../.;  
	   doCheck = false;
	   
		vendorHash = lib.fakeHash;
		
	   meta = with lib; {
		 description = "";
	   };
	 }
	