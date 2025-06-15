// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tiledMap, err := UnmarshalTiledMap(bytes)
//    bytes, err = tiledMap.Marshal()

package main

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalTiledMap(data []byte) (TiledMap, error) {
	var r TiledMap
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TiledMap) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Schema for the Tiled Map object.
type TiledMap struct {
	// Hex-formatted color (#RRGGBB or #AARRGGBB) (optional)                                                       
	Backgroundcolor                                                                            *string             `json:"backgroundcolor,omitempty"`
	// The class of the map (since 1.9, optional)                                                                  
	Class                                                                                      *string             `json:"class,omitempty"`
	// The compression level to use for tile layer data (defaults to -1, which means to use the                    
	// algorithm default)                                                                                          
	Compressionlevel                                                                           *int64              `json:"compressionlevel,omitempty"`
	// Number of tile rows                                                                                         
	Height                                                                                     int64               `json:"height"`
	// Length of the side of a hex tile in pixels (hexagonal maps only)                                            
	Hexsidelength                                                                              *int64              `json:"hexsidelength,omitempty"`
	// Whether the map has infinite dimensions                                                                     
	Infinite                                                                                   *bool               `json:"infinite,omitempty"`
	// Array of Layers                                                                                             
	Layers                                                                                     []TiledLayerElement `json:"layers"`
	// Auto-increments for each layer                                                                              
	Nextlayerid                                                                                int64               `json:"nextlayerid"`
	// Auto-increments for each placed object                                                                      
	Nextobjectid                                                                               int64               `json:"nextobjectid"`
	// `orthogonal`, `isometric`, `staggered` or `hexagonal`                                                       
	Orientation                                                                                TiledMapOrientation `json:"orientation"`
	// X coordinate of the parallax origin in pixels (since 1.8, default: 0)                                       
	Parallaxoriginx                                                                            *float64            `json:"parallaxoriginx,omitempty"`
	// Y coordinate of the parallax origin in pixels (since 1.8, default: 0)                                       
	Parallaxoriginy                                                                            *float64            `json:"parallaxoriginy,omitempty"`
	// Array of Properties                                                                                         
	Properties                                                                                 []TiledProperty     `json:"properties,omitempty"`
	// `right-down` (the default), `right-up`, `left-down` or `left-up` (currently only                            
	// supported for orthogonal maps)                                                                              
	Renderorder                                                                                Renderorder         `json:"renderorder"`
	// `x` or `y` (staggered / hexagonal maps only)                                                                
	Staggeraxis                                                                                *Staggeraxis        `json:"staggeraxis,omitempty"`
	// `odd` or `even` (staggered / hexagonal maps only)                                                           
	Staggerindex                                                                               *Staggerindex       `json:"staggerindex,omitempty"`
	// The Tiled version used to save the file                                                                     
	Tiledversion                                                                               *string             `json:"tiledversion,omitempty"`
	// Map grid height                                                                                             
	Tileheight                                                                                 int64               `json:"tileheight"`
	// Array of Tilesets                                                                                           
	Tilesets                                                                                   []TiledTileset      `json:"tilesets"`
	// Map grid width                                                                                              
	Tilewidth                                                                                  int64               `json:"tilewidth"`
	// `map` (since 1.0)                                                                                           
	Type                                                                                       TiledMapType        `json:"type"`
	// The JSON format version (previously a number, saved as string since 1.6)                                    
	Version                                                                                    string              `json:"version"`
	// Number of tile columns                                                                                      
	Width                                                                                      int64               `json:"width"`
}

// Schema for a Tiled Layer, which can be polymorphic based on its 'type' field.
type TiledLayerElement struct {
	// The class of the layer (since 1.9, optional)                                                                
	Class                                                                                    *string               `json:"class,omitempty"`
	// Incremental ID - unique across all layers                                                                   
	ID                                                                                       int64                 `json:"id"`
	// Whether layer is locked in the editor (default: false). (since 1.8.2)                                       
	Locked                                                                                   *bool                 `json:"locked,omitempty"`
	// Name assigned to this layer                                                                                 
	Name                                                                                     string                `json:"name"`
	// Horizontal layer offset in pixels (default: 0)                                                              
	Offsetx                                                                                  *float64              `json:"offsetx,omitempty"`
	// Vertical layer offset in pixels (default: 0)                                                                
	Offsety                                                                                  *float64              `json:"offsety,omitempty"`
	// Value between 0 and 1                                                                                       
	Opacity                                                                                  float64               `json:"opacity"`
	// Horizontal parallax factor for this layer (default: 1). (since 1.5)                                         
	Parallaxx                                                                                *float64              `json:"parallaxx,omitempty"`
	// Vertical parallax factor for this layer (default: 1). (since 1.5)                                           
	Parallaxy                                                                                *float64              `json:"parallaxy,omitempty"`
	// Array of Properties                                                                                         
	Properties                                                                               []TiledProperty       `json:"properties,omitempty"`
	// Hex-formatted tint color (#RRGGBB or #AARRGGBB) that is multiplied with any graphics                        
	// drawn by this layer or any child layers (optional).                                                         
	Tintcolor                                                                                *string               `json:"tintcolor,omitempty"`
	// `tilelayer`, `objectgroup`, `imagelayer` or `group`                                                         
	Type                                                                                     PurpleType            `json:"type"`
	// Whether layer is shown or hidden in editor                                                                  
	Visible                                                                                  bool                  `json:"visible"`
	// Horizontal layer offset in tiles. Always 0.                                                                 
	X                                                                                        int64                 `json:"x"`
	// Vertical layer offset in tiles. Always 0.                                                                   
	Y                                                                                        int64                 `json:"y"`
	// Array of chunks (optional). `tilelayer` only.                                                               
	Chunks                                                                                   []TiledTileLayerChunk `json:"chunks,omitempty"`
	// `zlib`, `gzip`, `zstd` (since 1.3) or empty (default). `tilelayer` only.                                    
	Compression                                                                              *Compression          `json:"compression,omitempty"`
	// Array of `unsigned int` (GIDs) or base64-encoded data. `tilelayer` only.                                    
	Data                                                                                     *Data                 `json:"data"`
	// `csv` (default) or `base64`. `tilelayer` only.                                                              
	Encoding                                                                                 *Encoding             `json:"encoding,omitempty"`
	// Row count. Same as map height for fixed-size maps. `tilelayer` only.                                        
	Height                                                                                   *int64                `json:"height,omitempty"`
	// X coordinate where layer content starts (for infinite maps)                                                 
	Startx                                                                                   *int64                `json:"startx,omitempty"`
	// Y coordinate where layer content starts (for infinite maps)                                                 
	Starty                                                                                   *int64                `json:"starty,omitempty"`
	// Column count. Same as map width for fixed-size maps. `tilelayer` only.                                      
	Width                                                                                    *int64                `json:"width,omitempty"`
	// `topdown` (default) or `index`. `objectgroup` only.                                                         
	Draworder                                                                                *Draworder            `json:"draworder,omitempty"`
	// Array of objects. `objectgroup` only.                                                                       
	Objects                                                                                  []TiledObject         `json:"objects,omitempty"`
	// Image used by this layer. `imagelayer` only.                                                                
	Image                                                                                    *string               `json:"image,omitempty"`
	// Height of the image used by this layer. `imagelayer` only. (since 1.11.1)                                   
	Imageheight                                                                              *int64                `json:"imageheight,omitempty"`
	// Width of the image used by this layer. `imagelayer` only. (since 1.11.1)                                    
	Imagewidth                                                                               *int64                `json:"imagewidth,omitempty"`
	// Whether the image drawn by this layer is repeated along the X axis. `imagelayer` only.                      
	// (since 1.8)                                                                                                 
	Repeatx                                                                                  *bool                 `json:"repeatx,omitempty"`
	// Whether the image drawn by this layer is repeated along the Y axis. `imagelayer` only.                      
	// (since 1.8)                                                                                                 
	Repeaty                                                                                  *bool                 `json:"repeaty,omitempty"`
	// Hex-formatted color (#RRGGBB) (optional). `imagelayer` only.                                                
	Transparentcolor                                                                         *string               `json:"transparentcolor,omitempty"`
	// Array of layers. `group` only.                                                                              
	Layers                                                                                   []TiledLayerElement   `json:"layers,omitempty"`
}

// Schema for a chunk of tile data in an infinite map. (Inferred from common patterns, not
// explicitly detailed in documentation snippets)
type TiledTileLayerChunk struct {
	// Array of GIDs for the chunk                    
	Data                                      []int64 `json:"data"`
	// Row count of the chunk                         
	Height                                    int64   `json:"height"`
	// Column count of the chunk                      
	Width                                     int64   `json:"width"`
	// X coordinate where chunk content starts        
	X                                         int64   `json:"x"`
	// Y coordinate where chunk content starts        
	Y                                         int64   `json:"y"`
}

// Schema for an Object on an object layer.
type TiledObject struct {
	// Used to mark an object as an ellipse                                                
	Ellipse                                                               *bool            `json:"ellipse,omitempty"`
	// Global tile ID, only if object represents a tile                                    
	Gid                                                                   *int64           `json:"gid,omitempty"`
	// Height in pixels.                                                                   
	Height                                                                float64          `json:"height"`
	// Incremental ID, unique across all objects                                           
	ID                                                                    int64            `json:"id"`
	// String assigned to name field in editor                                             
	Name                                                                  string           `json:"name"`
	// Used to mark an object as a point                                                   
	Point                                                                 *bool            `json:"point,omitempty"`
	// Array of Points, in case the object is a polygon                                    
	Polygon                                                               []TiledPoint     `json:"polygon,omitempty"`
	// Array of Points, in case the object is a polyline                                   
	Polyline                                                              []TiledPoint     `json:"polyline,omitempty"`
	// Array of Properties                                                                 
	Properties                                                            []TiledProperty  `json:"properties,omitempty"`
	// Angle in degrees clockwise                                                          
	Rotation                                                              *float64         `json:"rotation,omitempty"`
	// Reference to a template file, in case object is a template instance                 
	Template                                                              *string          `json:"template,omitempty"`
	// Only used for text objects                                                          
	Text                                                                  *TiledTextObject `json:"text,omitempty"`
	// The class of the object (was saved as `class` in 1.9, optional)                     
	Type                                                                  *string          `json:"type,omitempty"`
	// Whether object is shown in editor.                                                  
	Visible                                                               bool             `json:"visible"`
	// Width in pixels.                                                                    
	Width                                                                 float64          `json:"width"`
	// X coordinate in pixels                                                              
	X                                                                     float64          `json:"x"`
	// Y coordinate in pixels                                                              
	Y                                                                     float64          `json:"y"`
}

// Schema for a 2D coordinate point (inferred from common conventions).
type TiledPoint struct {
	// X coordinate        
	X              float64 `json:"x"`
	// Y coordinate        
	Y              float64 `json:"y"`
}

// Schema for a custom Property.
type TiledProperty struct {
	// Name of the property                                                                                       
	Name                                                                                        string            `json:"name"`
	// Name of the custom property type, when applicable (since 1.8)                                              
	Propertytype                                                                                *string           `json:"propertytype,omitempty"`
	// Type of the property (`string` (default), `int`, `float`, `bool`, `color`, `file`,                         
	// `object` or `class` (since 0.16, with `color` and `file` added in 0.17, `object` added in                  
	// 1.4 and `class` added in 1.8))                                                                             
	Type                                                                                        TiledPropertyType `json:"type"`
	// Value of the property                                                                                      
	Value                                                                                       *Value            `json:"value"`
}

// Only used for text objects
//
// Schema for a Text object within Tiled.
type TiledTextObject struct {
	// Whether to use a bold font (default: `false`)                                  
	Bold                                                                      *bool   `json:"bold,omitempty"`
	// Hex-formatted color (#RRGGBB or #AARRGGBB) (default: `#000000`)                
	Color                                                                     *string `json:"color,omitempty"`
	// Font family (default: `sans-serif`)                                            
	Fontfamily                                                                *string `json:"fontfamily,omitempty"`
	// Horizontal alignment (`center`, `right`, `justify` or `left` (default))        
	Halign                                                                    *Halign `json:"halign,omitempty"`
	// Whether to use an italic font (default: `false`)                               
	Italic                                                                    *bool   `json:"italic,omitempty"`
	// Whether to use kerning when placing characters (default: `true`)               
	Kerning                                                                   *bool   `json:"kerning,omitempty"`
	// Pixel size of font (default: 16)                                               
	Pixelsize                                                                 *int64  `json:"pixelsize,omitempty"`
	// Whether to strike out the text (default: `false`)                              
	Strikeout                                                                 *bool   `json:"strikeout,omitempty"`
	// Text                                                                           
	Text                                                                      string  `json:"text"`
	// Whether to underline the text (default: `false`)                               
	Underline                                                                 *bool   `json:"underline,omitempty"`
	// Vertical alignment (`center`, `bottom` or `top` (default))                     
	Valign                                                                    *Valign `json:"valign,omitempty"`
	// Whether the text is wrapped within the object bounds (default: `false`)        
	Wrap                                                                      *bool   `json:"wrap,omitempty"`
}

// Schema for a Tiled Tileset.
type TiledTileset struct {
	// Hex-formatted color (#RRGGBB or #AARRGGBB) (optional)                                                                 
	Backgroundcolor                                                                             *string                      `json:"backgroundcolor,omitempty"`
	// The class of the tileset (since 1.9, optional)                                                                        
	Class                                                                                       *string                      `json:"class,omitempty"`
	// The number of tile columns in the tileset                                                                             
	Columns                                                                                     int64                        `json:"columns"`
	// The fill mode to use when rendering tiles from this tileset (`stretch` (default) or                                   
	// `preserve-aspect-fit`) (since 1.9)                                                                                    
	Fillmode                                                                                    *Fillmode                    `json:"fillmode,omitempty"`
	// GID corresponding to the first tile in the set                                                                        
	Firstgid                                                                                    int64                        `json:"firstgid"`
	// (optional)                                                                                                            
	Grid                                                                                        *TiledTilesetGrid            `json:"grid,omitempty"`
	// Image used for tiles in this set                                                                                      
	Image                                                                                       *string                      `json:"image,omitempty"`
	// Height of source image in pixels                                                                                      
	Imageheight                                                                                 *int64                       `json:"imageheight,omitempty"`
	// Width of source image in pixels                                                                                       
	Imagewidth                                                                                  *int64                       `json:"imagewidth,omitempty"`
	// Buffer between image edge and first tile (pixels)                                                                     
	Margin                                                                                      *int64                       `json:"margin,omitempty"`
	// Name given to this tileset                                                                                            
	Name                                                                                        string                       `json:"name"`
	// Alignment to use for tile objects (`unspecified` (default), `topleft`, `top`, `topright`,                             
	// `left`, `center`, `right`, `bottomleft`, `bottom` or `bottomright`) (since 1.4)                                       
	Objectalignment                                                                             *Objectalignment             `json:"objectalignment,omitempty"`
	// Array of Properties                                                                                                   
	Properties                                                                                  []TiledProperty              `json:"properties,omitempty"`
	// The external file that contains this tilesets data                                                                    
	Source                                                                                      *string                      `json:"source,omitempty"`
	// Spacing between adjacent tiles in image (pixels)                                                                      
	Spacing                                                                                     *int64                       `json:"spacing,omitempty"`
	// Array of Terrains (optional)                                                                                          
	Terrains                                                                                    []TiledTerrain               `json:"terrains,omitempty"`
	// The number of tiles in this tileset                                                                                   
	Tilecount                                                                                   int64                        `json:"tilecount"`
	// The Tiled version used to save the file                                                                               
	Tiledversion                                                                                *string                      `json:"tiledversion,omitempty"`
	// Maximum height of tiles in this set                                                                                   
	Tileheight                                                                                  int64                        `json:"tileheight"`
	// (optional)                                                                                                            
	Tileoffset                                                                                  *TiledTileOffset             `json:"tileoffset,omitempty"`
	// The size to use when rendering tiles from this tileset on a tile layer (`tile` (default)                              
	// or `grid`) (since 1.9)                                                                                                
	Tilerendersize                                                                              *Tilerendersize              `json:"tilerendersize,omitempty"`
	// Array of Tiles (optional)                                                                                             
	Tiles                                                                                       []TiledTileDefinition        `json:"tiles,omitempty"`
	// Maximum width of tiles in this set                                                                                    
	Tilewidth                                                                                   int64                        `json:"tilewidth"`
	// Allowed transformations (optional)                                                                                    
	Transformations                                                                             *TiledTilesetTransformations `json:"transformations,omitempty"`
	// Hex-formatted color (#RRGGBB) (optional)                                                                              
	Transparentcolor                                                                            *string                      `json:"transparentcolor,omitempty"`
	// `tileset` (for tileset files, since 1.0)                                                                              
	Type                                                                                        TilesetType                  `json:"type"`
	// The JSON format version (previously a number, saved as string since 1.6)                                              
	Version                                                                                     string                       `json:"version"`
	// Array of Wang sets (since 1.1.5)                                                                                      
	Wangsets                                                                                    []TiledWangSet               `json:"wangsets,omitempty"`
}

// (optional)
//
// Schema for common grid settings for tiles within a tileset.
type TiledTilesetGrid struct {
	// Cell height of tile grid                             
	Height                                  int64           `json:"height"`
	// `orthogonal` (default) or `isometric`                
	Orientation                             GridOrientation `json:"orientation"`
	// Cell width of tile grid                              
	Width                                   int64           `json:"width"`
}

// Schema for a terrain type within a tileset.
type TiledTerrain struct {
	// Name of terrain                                      
	Name                                    string          `json:"name"`
	// Array of Properties                                  
	Properties                              []TiledProperty `json:"properties,omitempty"`
	// Local ID of tile representing terrain                
	Tile                                    int64           `json:"tile"`
}

// (optional)
//
// Schema for an offset in pixels for drawing tiles from a tileset.
type TiledTileOffset struct {
	// Horizontal offset in pixels                       
	X                                              int64 `json:"x"`
	// Vertical offset in pixels (positive is down)      
	Y                                              int64 `json:"y"`
}

// Schema for an individual Tile within a Tileset.
type TiledTileDefinition struct {
	// Array of Frames for tile animation                                                                             
	Animation                                                                                   []TiledAnimationFrame `json:"animation,omitempty"`
	// The height of the sub-rectangle representing this tile (defaults to the image height)                          
	Height                                                                                      *int64                `json:"height,omitempty"`
	// Local ID of the tile                                                                                           
	ID                                                                                          int64                 `json:"id"`
	// Image representing this tile (optional, used for image collection tilesets)                                    
	Image                                                                                       *string               `json:"image,omitempty"`
	// Height of the tile image in pixels                                                                             
	Imageheight                                                                                 *int64                `json:"imageheight,omitempty"`
	// Width of the tile image in pixels                                                                              
	Imagewidth                                                                                  *int64                `json:"imagewidth,omitempty"`
	// Layer with type `objectgroup`, when collision shapes are specified (optional)                                  
	Objectgroup                                                                                 *TileTiledLayer       `json:"objectgroup,omitempty"`
	// Percentage chance this tile is chosen when competing with others in the editor (optional)                      
	Probability                                                                                 *float64              `json:"probability,omitempty"`
	// Array of Properties                                                                                            
	Properties                                                                                  []TiledProperty       `json:"properties,omitempty"`
	// Index of terrain for each corner of tile (optional, replaced by Wang sets since 1.5)                           
	Terrain                                                                                     []int64               `json:"terrain,omitempty"`
	// The class of the tile (was saved as `class` in 1.9, optional)                                                  
	Type                                                                                        *string               `json:"type,omitempty"`
	// The width of the sub-rectangle representing this tile (defaults to the image width)                            
	Width                                                                                       *int64                `json:"width,omitempty"`
	// The X position of the sub-rectangle representing this tile (default: 0)                                        
	X                                                                                           *int64                `json:"x,omitempty"`
	// The Y position of the sub-rectangle representing this tile (default: 0)                                        
	Y                                                                                           *int64                `json:"y,omitempty"`
}

// Schema for a single frame in a tile animation.
type TiledAnimationFrame struct {
	// Frame duration in milliseconds             
	Duration                                int64 `json:"duration"`
	// Local tile ID representing this frame      
	Tileid                                  int64 `json:"tileid"`
}

// Layer with type `objectgroup`, when collision shapes are specified (optional)
//
// Schema for a Tiled Layer, which can be polymorphic based on its 'type' field.
type TileTiledLayer struct {
	// `tilelayer`, `objectgroup`, `imagelayer` or `group`                                                         
	Type                                                                                     FluffyType            `json:"type"`
	// The class of the layer (since 1.9, optional)                                                                
	Class                                                                                    *string               `json:"class,omitempty"`
	// Incremental ID - unique across all layers                                                                   
	ID                                                                                       int64                 `json:"id"`
	// Whether layer is locked in the editor (default: false). (since 1.8.2)                                       
	Locked                                                                                   *bool                 `json:"locked,omitempty"`
	// Name assigned to this layer                                                                                 
	Name                                                                                     string                `json:"name"`
	// Horizontal layer offset in pixels (default: 0)                                                              
	Offsetx                                                                                  *float64              `json:"offsetx,omitempty"`
	// Vertical layer offset in pixels (default: 0)                                                                
	Offsety                                                                                  *float64              `json:"offsety,omitempty"`
	// Value between 0 and 1                                                                                       
	Opacity                                                                                  float64               `json:"opacity"`
	// Horizontal parallax factor for this layer (default: 1). (since 1.5)                                         
	Parallaxx                                                                                *float64              `json:"parallaxx,omitempty"`
	// Vertical parallax factor for this layer (default: 1). (since 1.5)                                           
	Parallaxy                                                                                *float64              `json:"parallaxy,omitempty"`
	// Array of Properties                                                                                         
	Properties                                                                               []TiledProperty       `json:"properties,omitempty"`
	// Hex-formatted tint color (#RRGGBB or #AARRGGBB) that is multiplied with any graphics                        
	// drawn by this layer or any child layers (optional).                                                         
	Tintcolor                                                                                *string               `json:"tintcolor,omitempty"`
	// Whether layer is shown or hidden in editor                                                                  
	Visible                                                                                  bool                  `json:"visible"`
	// Horizontal layer offset in tiles. Always 0.                                                                 
	X                                                                                        int64                 `json:"x"`
	// Vertical layer offset in tiles. Always 0.                                                                   
	Y                                                                                        int64                 `json:"y"`
	// Array of chunks (optional). `tilelayer` only.                                                               
	Chunks                                                                                   []TiledTileLayerChunk `json:"chunks,omitempty"`
	// `zlib`, `gzip`, `zstd` (since 1.3) or empty (default). `tilelayer` only.                                    
	Compression                                                                              *Compression          `json:"compression,omitempty"`
	// Array of `unsigned int` (GIDs) or base64-encoded data. `tilelayer` only.                                    
	Data                                                                                     *Data                 `json:"data"`
	// `csv` (default) or `base64`. `tilelayer` only.                                                              
	Encoding                                                                                 *Encoding             `json:"encoding,omitempty"`
	// Row count. Same as map height for fixed-size maps. `tilelayer` only.                                        
	Height                                                                                   *int64                `json:"height,omitempty"`
	// X coordinate where layer content starts (for infinite maps)                                                 
	Startx                                                                                   *int64                `json:"startx,omitempty"`
	// Y coordinate where layer content starts (for infinite maps)                                                 
	Starty                                                                                   *int64                `json:"starty,omitempty"`
	// Column count. Same as map width for fixed-size maps. `tilelayer` only.                                      
	Width                                                                                    *int64                `json:"width,omitempty"`
	// `topdown` (default) or `index`. `objectgroup` only.                                                         
	Draworder                                                                                *Draworder            `json:"draworder,omitempty"`
	// Array of objects. `objectgroup` only.                                                                       
	Objects                                                                                  []TiledObject         `json:"objects,omitempty"`
	// Image used by this layer. `imagelayer` only.                                                                
	Image                                                                                    *string               `json:"image,omitempty"`
	// Height of the image used by this layer. `imagelayer` only. (since 1.11.1)                                   
	Imageheight                                                                              *int64                `json:"imageheight,omitempty"`
	// Width of the image used by this layer. `imagelayer` only. (since 1.11.1)                                    
	Imagewidth                                                                               *int64                `json:"imagewidth,omitempty"`
	// Whether the image drawn by this layer is repeated along the X axis. `imagelayer` only.                      
	// (since 1.8)                                                                                                 
	Repeatx                                                                                  *bool                 `json:"repeatx,omitempty"`
	// Whether the image drawn by this layer is repeated along the Y axis. `imagelayer` only.                      
	// (since 1.8)                                                                                                 
	Repeaty                                                                                  *bool                 `json:"repeaty,omitempty"`
	// Hex-formatted color (#RRGGBB) (optional). `imagelayer` only.                                                
	Transparentcolor                                                                         *string               `json:"transparentcolor,omitempty"`
	// Array of layers. `group` only.                                                                              
	Layers                                                                                   []TiledLayerElement   `json:"layers,omitempty"`
}

// Allowed transformations (optional)
//
// Schema for allowed transformations for tiles within a tileset.
type TiledTilesetTransformations struct {
	// Indicates whether tiles can be flipped horizontally.                                         
	Hflip                                                                                      bool `json:"hflip"`
	// Indicates whether untransformed tiles remain preferred, or if transformed tiles are used     
	// to produce more variations.                                                                  
	Preferuntransformed                                                                        bool `json:"preferuntransformed"`
	// Indicates whether tiles can be rotated in 90-degree increments.                              
	Rotate                                                                                     bool `json:"rotate"`
	// Indicates whether tiles can be flipped vertically.                                           
	Vflip                                                                                      bool `json:"vflip"`
}

// Schema for a Wang set for advanced tile auto-mapping.
type TiledWangSet struct {
	// The class of the Wang set (since Tiled 1.9, optional).                 
	Class                                                    *string          `json:"class,omitempty"`
	// An array of Wang colors (since Tiled 1.5).                             
	Colors                                                   []TiledWangColor `json:"colors"`
	// The name of the Wang set.                                              
	Name                                                     string           `json:"name"`
	// An array of Properties.                                                
	Properties                                               []TiledProperty  `json:"properties,omitempty"`
	// The local ID of the tile representing the Wang set.                    
	Tile                                                     int64            `json:"tile"`
	// Can be `corner`, `edge`, or `mixed` (since Tiled 1.5).                 
	Type                                                     WangsetType      `json:"type"`
	// An array of Wang tiles.                                                
	Wangtiles                                                []TiledWangTile  `json:"wangtiles"`
}

// Schema for a color used in Wang sets (for corners or edges).
type TiledWangColor struct {
	// The class of the Wang color (since 1.9, optional)                
	Class                                               *string         `json:"class,omitempty"`
	// Hex-formatted color (#RRGGBB or #AARRGGBB)                       
	Color                                               string          `json:"color"`
	// Name of the Wang color                                           
	Name                                                string          `json:"name"`
	// Probability used when randomizing                                
	Probability                                         float64         `json:"probability"`
	// Array of Properties (since 1.5)                                  
	Properties                                          []TiledProperty `json:"properties,omitempty"`
	// Local ID of tile representing the Wang color                     
	Tile                                                int64           `json:"tile"`
}

// Schema for a specific tile's Wang ID (color indexes for its corners/edges).
type TiledWangTile struct {
	// Local ID of the tile.                         
	Tileid                                   int64   `json:"tileid"`
	// Array of Wang color indexes (`uchar`).        
	Wangid                                   []int64 `json:"wangid"`
}

// `zlib`, `gzip`, `zstd` (since 1.3) or empty (default). `tilelayer` only.
type Compression string

const (
	Empty Compression = ""
	Gzip  Compression = "gzip"
	Zlib  Compression = "zlib"
	Zstd  Compression = "zstd"
)

// `topdown` (default) or `index`. `objectgroup` only.
type Draworder string

const (
	Index   Draworder = "index"
	Topdown Draworder = "topdown"
)

// `csv` (default) or `base64`. `tilelayer` only.
type Encoding string

const (
	Base64 Encoding = "base64"
	CSV    Encoding = "csv"
)

// Type of the property (`string` (default), `int`, `float`, `bool`, `color`, `file`,
// `object` or `class` (since 0.16, with `color` and `file` added in 0.17, `object` added in
// 1.4 and `class` added in 1.8))
type TiledPropertyType string

const (
	Bool   TiledPropertyType = "bool"
	Class  TiledPropertyType = "class"
	Color  TiledPropertyType = "color"
	File   TiledPropertyType = "file"
	Float  TiledPropertyType = "float"
	Int    TiledPropertyType = "int"
	Object TiledPropertyType = "object"
	String TiledPropertyType = "string"
)

// Horizontal alignment (`center`, `right`, `justify` or `left` (default))
type Halign string

const (
	HalignCenter Halign = "center"
	HalignLeft   Halign = "left"
	HalignRight  Halign = "right"
	Justify      Halign = "justify"
)

// Vertical alignment (`center`, `bottom` or `top` (default))
type Valign string

const (
	ValignBottom Valign = "bottom"
	ValignCenter Valign = "center"
	ValignTop    Valign = "top"
)

// `tilelayer`, `objectgroup`, `imagelayer` or `group`
type PurpleType string

const (
	Group             PurpleType = "group"
	Imagelayer        PurpleType = "imagelayer"
	PurpleObjectgroup PurpleType = "objectgroup"
	Tilelayer         PurpleType = "tilelayer"
)

// `orthogonal`, `isometric`, `staggered` or `hexagonal`
type TiledMapOrientation string

const (
	Hexagonal        TiledMapOrientation = "hexagonal"
	PurpleIsometric  TiledMapOrientation = "isometric"
	PurpleOrthogonal TiledMapOrientation = "orthogonal"
	Staggered        TiledMapOrientation = "staggered"
)

// `right-down` (the default), `right-up`, `left-down` or `left-up` (currently only
// supported for orthogonal maps)
type Renderorder string

const (
	LeftDown  Renderorder = "left-down"
	LeftUp    Renderorder = "left-up"
	RightDown Renderorder = "right-down"
	RightUp   Renderorder = "right-up"
)

// `x` or `y` (staggered / hexagonal maps only)
type Staggeraxis string

const (
	X Staggeraxis = "x"
	Y Staggeraxis = "y"
)

// `odd` or `even` (staggered / hexagonal maps only)
type Staggerindex string

const (
	Even Staggerindex = "even"
	Odd  Staggerindex = "odd"
)

// The fill mode to use when rendering tiles from this tileset (`stretch` (default) or
// `preserve-aspect-fit`) (since 1.9)
type Fillmode string

const (
	PreserveAspectFit Fillmode = "preserve-aspect-fit"
	Stretch           Fillmode = "stretch"
)

// `orthogonal` (default) or `isometric`
type GridOrientation string

const (
	FluffyIsometric  GridOrientation = "isometric"
	FluffyOrthogonal GridOrientation = "orthogonal"
)

// Alignment to use for tile objects (`unspecified` (default), `topleft`, `top`, `topright`,
// `left`, `center`, `right`, `bottomleft`, `bottom` or `bottomright`) (since 1.4)
type Objectalignment string

const (
	Bottomleft            Objectalignment = "bottomleft"
	Bottomright           Objectalignment = "bottomright"
	ObjectalignmentBottom Objectalignment = "bottom"
	ObjectalignmentCenter Objectalignment = "center"
	ObjectalignmentLeft   Objectalignment = "left"
	ObjectalignmentRight  Objectalignment = "right"
	ObjectalignmentTop    Objectalignment = "top"
	Topleft               Objectalignment = "topleft"
	Topright              Objectalignment = "topright"
	Unspecified           Objectalignment = "unspecified"
)

// The size to use when rendering tiles from this tileset on a tile layer (`tile` (default)
// or `grid`) (since 1.9)
type Tilerendersize string

const (
	Grid Tilerendersize = "grid"
	Tile Tilerendersize = "tile"
)

// `tilelayer`, `objectgroup`, `imagelayer` or `group`
type FluffyType string

const (
	FluffyObjectgroup FluffyType = "objectgroup"
)

type TilesetType string

const (
	Tileset TilesetType = "tileset"
)

// Can be `corner`, `edge`, or `mixed` (since Tiled 1.5).
type WangsetType string

const (
	Corner WangsetType = "corner"
	Edge   WangsetType = "edge"
	Mixed  WangsetType = "mixed"
)

type TiledMapType string

const (
	Map TiledMapType = "map"
)

// Array of `unsigned int` (GIDs) or base64-encoded data. `tilelayer` only.
type Data struct {
	IntegerArray []int64
	String       *string
}

func (x *Data) UnmarshalJSON(data []byte) error {
	x.IntegerArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.IntegerArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Data) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.IntegerArray != nil, x.IntegerArray, false, nil, false, nil, false, nil, false)
}

type Value struct {
	AnythingMap map[string]interface{}
	Bool        *bool
	Double      *float64
	Integer     *int64
	String      *string
}

func (x *Value) UnmarshalJSON(data []byte) error {
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, false, nil, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Value) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, false, nil, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
