package image_file_local_file_repository

import "wallshrink/domain"

func (r *imageFileLocalFileRepository) LoadImageFile(filePath string) (domain.ImageFile, error) {
	stem, extension := splitFileName(filePath)

	imageFile := domain.ImageFile{
		Stem:      stem,
		Extension: extension,
	}

	// Get file size
	size, err := getFileSize(filePath)
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Size = size

	// Get image dimension
	dimension, err := getImageDimension(filePath)
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Dimension = dimension

	return imageFile, nil
}