using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.MicroserviceDtos;
using Explorer.Tours.API.Public;
using Explorer.Tours.API.Public.Administration;
using Explorer.Tours.Core.UseCases;
using Explorer.Tours.Core.UseCases.Administration;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.Net.Http;
using System.Text;
using System.Text.Json;

namespace Explorer.API.Controllers.Author
{
    [Authorize(Policy = "authorPolicy")]
    [Route("api/tourKeyPoint")]
    public class TourKeyPointController : BaseApiController
    {
        private readonly ITourKeyPointService _tourKeyPointService;
        private readonly IPublicTourKeyPointService _publicTourKeyPointService;
        private readonly IHttpClientFactory _factory;

        public TourKeyPointController(ITourKeyPointService tourKeyPointService, IPublicTourKeyPointService publicTourKeyPointService, IHttpClientFactory factory)
        {
            _tourKeyPointService = tourKeyPointService;
            _publicTourKeyPointService = publicTourKeyPointService;
            _factory = factory;
        }

        [HttpGet]
        public ActionResult<PagedResult<TourKeyPointDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourKeyPointService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        [HttpGet("tour/{tourId:int}")]
        public ActionResult<PagedResult<TourKeyPointDto>> GetByTourId(int tourId)
        {
            var result = _tourKeyPointService.GetByTourId(tourId);
            return CreateResponse(result);
        }

        [HttpGet("{id:Guid}")]
        public async Task<ActionResult<TourKeypointDto>> Get(Guid id)
        {
            //var result = _tourKeyPointService.Get(id);

            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.GetAsync("tourKeypoints/" + id.ToString());
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            TourKeypointDto tourKeyPointDto =
                JsonSerializer.Deserialize<TourKeypointDto>(jsonResponse);

            return Ok(tourKeyPointDto);
        }

        [HttpPost]
        public async Task<ActionResult> Create([FromBody] TourKeypointDto tourKeyPoint)
        {
            //var result = _tourKeyPointService.Create(tourKeyPoint);

            var client = _factory.CreateClient("toursMicroservice");
            using StringContent jsonContent = new(
                JsonSerializer.Serialize(tourKeyPoint),
                Encoding.UTF8,
                "application/json");

            using HttpResponseMessage response = await client.PostAsync(
                "/tourKeypoints",
                jsonContent);

            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }


            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }



        [HttpPut("{id:int}")]
        public ActionResult<TourKeyPointDto> Update([FromBody] TourKeyPointDto tourKeyPoint)
        {
            var result = _tourKeyPointService.Update(tourKeyPoint);
            return CreateResponse(result);
        }

        [HttpDelete("{id:int}")]
        public ActionResult Delete(int id)
        {
            var result = _tourKeyPointService.Delete(id);
            return CreateResponse(result);

        }



        [HttpGet("public")]
        public ActionResult<PagedResult<PublicTourKeyPointDto>> GetAllPublic([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _publicTourKeyPointService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        [HttpPost("public")]
        public ActionResult<PublicTourKeyPointDto> CreatePublic([FromBody] PublicTourKeyPointDto tourKeyPoint)
        {
            var result = _publicTourKeyPointService.Create(tourKeyPoint);
            return CreateResponse(result);
        }

        [HttpPut("public/{tourId}/{status}")]
        public ActionResult<PublicTourKeyPointDto> ChangeStatus(int tourId, String status)
        {
            var result = _publicTourKeyPointService.ChangeStatus(tourId, status);
            return CreateResponse(result);
        }

        [HttpGet("public/{status}")]
        public ActionResult<PagedResult<PublicTourKeyPointDto>> GetByStatus(String status)
        {
            var result = _publicTourKeyPointService.GetByStatus(status);
            return CreateResponse(result);
        }

    }
}
