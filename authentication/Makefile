mockgen_task:
	$(eval SERVICE_LIST := $(call get_service_list))	
	$(foreach file, $(SERVICE_LIST), $(call mockgen_service,$(shell basename $(file))))
	$(eval REPOSITORY_LIST := $(call get_repository_list))
	$(foreach file, $(REPOSITORY_LIST), $(call mockgen_repository,$(shell basename $(file))))

define get_service_list
	$(shell	find ./service/ -maxdepth 1 -type f ! -name "*_impl*.go")
endef

define mockgen_service
	$(shell mockgen -source ./service/$1 -destination ./service/mock/$1)
endef

define get_repository_list
	$(shell	find ./domain/repository -maxdepth 1 -type f )
endef

define mockgen_repository
	$(shell mockgen -source ./domain/repository/$1 -destination ./domain/repository/mock/$1)
endef