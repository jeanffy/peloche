package ui

import "peloche/internal/domain"

// ---------------------------------------------------------------------------
// #region root folder

const EventRootFolderSelected = "event.rootFolderSelected"

type EventRootFolderSelectedParams struct {
	RootFolderPath string
}

const EventRootFolderChanging = "event.rootFolderChanging"

type EventRootFolderChangingParams struct {
	RootFolderPath string
}

const EventRootFolderChanged = "event.rootFolderChanged"

type EventRootFolderChangedParams struct {
	RootFolderPath string
	FolderTree     *domain.FolderTree
}

// #endregion

// ---------------------------------------------------------------------------
// #region current folder

const EventCurrentFolderSelected = "event.currentFolderSelected"

type EventCurrentFolderSelectedParams struct {
	CurrentFolderPath string
}

const EventCurrentFolderChanging = "event.currentFolderChanging"

type EventCurrentFolderChangingParams struct {
	CurrentFolderPath string
}

const EventCurrentFolderChanged = "event.currentFolderChanged"

type EventCurrentFolderChangedParams struct {
	CurrentFolderPath string
	PhotoList         *domain.PhotoList
}

// #endregion

// ---------------------------------------------------------------------------
// #region user interaction

const EventThumbnailSizeChanged = "event.thumbnailSizeChanged"

type EventThumbnailSizeChangedParams struct {
	Size uint
}

const EventSelectedPhotoChanged = "event.selectedPhotoChanged"

type EventSelectedPhotoChangedParams struct {
	Index int
}

// #endregion
