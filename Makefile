bld:
	cat ./configs/.project_name | awk '{print $1}' | xargs bash ./scripts/build.sh "build" $1

clean:
	bash ./scripts/clean.sh "build"
