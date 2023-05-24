#!/usr/bin/env bash



main(){
	clean
	generate
	load
}	


load() {
	echo "load"
	for f in dest/*.yaml; do kubectl apply -f $f; done
}

clean() {
	rm -rf ./dest/*
}

generate () {

	read -p "number_of_applications: " number_of_applications
	read -p "base_of_applications: " base_of_applications
	for(( i=0; i<$number_of_applications; i++ ))
	do
		APP_NUMBER=$(($base_of_applications + $i))
		WAVE_LEVEL=$(($i - $base_of_applications))
		FULLNAME=httpbin-$APP_NUMBER
		echo $FULLNAME
		cp ./template/httpbin-application-base.yaml ./dest/$FULLNAME.yaml
		find ./dest/$FULLNAME.yaml -type f -exec sed -i '' -e 's/fullname/'"$FULLNAME"'/g' {} \;
	done
}


main @